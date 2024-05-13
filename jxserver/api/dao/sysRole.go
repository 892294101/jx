// 角色 数据层
// author xiaoRui

package dao

import (
	"github.com/jx/jxserver/api/entity"
	"github.com/jx/jxserver/common/util"
	. "github.com/jx/jxserver/pkg/db"
	"time"
)

// 根据名称查询
func GetSysRoleByName(roleName string) (sysRole entity.SysRole) {
	Db.Where("role_name = ?", roleName).First(&sysRole)
	return sysRole
}

// 根据角色Key查询
func GetSysRoleByKey(roleKey string) (sysRole entity.SysRole) {
	Db.Where("role_key = ?", roleKey).First(&sysRole)
	return sysRole
}

// 新增角色
func CreateSysRole(dto entity.AddSysRoleDto) bool {
	sysRoleByName := GetSysRoleByName(dto.RoleName)
	if sysRoleByName.ID > 0 {
		return false
	}
	sysRoleByKey := GetSysRoleByKey(dto.RoleKey)
	if sysRoleByKey.ID > 0 {
		return false
	}
	addSysRole := entity.SysRole{
		RoleName:    dto.RoleName,
		RoleKey:     dto.RoleKey,
		Description: dto.Description,
		Status:      dto.Status,
		CreateTime:  util.HTime{Time: time.Now()},
	}
	tx := Db.Create(&addSysRole)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// 根据id获取详情
func GetSysRoleById(Id int) (sysRole entity.SysRole) {
	Db.First(&sysRole, Id)
	return sysRole
}

// 修改角色
func UpdateSysRole(dto entity.UpdateSysRoleDto) (sysRole entity.SysRole) {
	Db.First(&sysRole, dto.Id)
	sysRole.RoleName = dto.RoleName
	sysRole.RoleKey = dto.RoleKey
	sysRole.Status = dto.Status
	if dto.Description != "" {
		sysRole.Description = dto.Description
	}
	Db.Save(&sysRole)
	return sysRole
}

// 根据id删除角色
func DeleteSysRoleById(dto entity.SysRoleIdDto) {
	Db.Table("ss_basicmanage_roles").Delete(&entity.SysRole{}, dto.Id)
	Db.Table("ss_basicmanage_roles_perms").Where("role_id = ?", dto.Id).Delete(&entity.SysRoleMenu{})
}

// 角色状态启用/停用
func UpdateSysRoleStatus(dto entity.UpdateSysRoleStatusDto) bool {
	var sysRole entity.SysRole
	Db.First(&sysRole, dto.Id)
	sysRole.Status = dto.Status
	tx := Db.Save(&sysRole)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// 分页查询角色列表
func GetSysRoleList(PageNum, PageSize int, RoleName, status, BeginTime, EndTime string) (sysRole []*entity.SysRole, count int64) {
	curDb := Db.Table("ss_basicmanage_roles")
	if RoleName != "" {
		curDb = curDb.Where("role_name = ?", RoleName)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if status != "" {
		curDb = curDb.Where("status = ?", status)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time DESC").Find(&sysRole)
	return sysRole, count
}

// 角色下拉列表
func QuerySysRoleVoList() (sysRoleVo []entity.SysRoleVo) {
	Db.Table("ss_basicmanage_roles").Select("id, role_name").Scan(&sysRoleVo)
	return sysRoleVo
}

// 根据角色的id查询菜单权限数据列表
func QueryRoleMenuIdList(Id int) (idVo []entity.IdVo) {
	const menuType int = 3
	Db.Table("ss_basicmanage_menu sm").
		Select("sm.id").
		Joins("LEFT JOIN ss_basicmanage_roles_perms srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN ss_basicmanage_roles sr ON sr.id = srm.role_id").
		Where("sm.menu_type = ?", menuType).
		Where("sr.id = ?", Id).
		Scan(&idVo)
	return idVo
}

// 分配权限
func AssignPermissions(menu entity.RoleMenu) (err error) {
	err = Db.Table("ss_basicmanage_roles_perms").Where("role_id = ?", menu.Id).Delete(&entity.SysRoleMenu{}).Error
	if err != nil {
		return err
	}
	for _, value := range menu.MenuIds {
		var entity entity.SysRoleMenu
		entity.RoleId = menu.Id
		entity.MenuId = value
		Db.Create(&entity)
	}
	return err
}
