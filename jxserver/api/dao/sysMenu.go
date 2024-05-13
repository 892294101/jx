// 菜单 数据层
// author xiaoRui

package dao

import (
	"github.com/jx/jxserver/api/entity"
	"github.com/jx/jxserver/common/util"
	. "github.com/jx/jxserver/pkg/db"
	"time"
)

// 根据菜单名称查询
func GetSysMenuByName(menuName string) (sysMenu entity.SysMenu) {
	Db.Where("menu_name = ?", menuName).First(&sysMenu)
	return sysMenu
}

// 新增菜单
func CreateSysMenu(addSysMenu entity.SysMenu) bool {
	sysMenuByName := GetSysMenuByName(addSysMenu.MenuName)
	if sysMenuByName.ID != 0 {
		return false
	}
	// 目录
	if addSysMenu.MenuType == 1 {
		sysMenu := entity.SysMenu{
			ParentId:   0,
			MenuName:   addSysMenu.MenuName,
			Icon:       addSysMenu.Icon,
			MenuType:   addSysMenu.MenuType,
			Url:        addSysMenu.Url,
			MenuStatus: addSysMenu.MenuStatus,
			Sort:       addSysMenu.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	} else if addSysMenu.MenuType == 2 {
		sysMenu := entity.SysMenu{
			ParentId:   addSysMenu.ParentId,
			MenuName:   addSysMenu.MenuName,
			Icon:       addSysMenu.Icon,
			MenuType:   addSysMenu.MenuType,
			MenuStatus: addSysMenu.MenuStatus,
			Value:      addSysMenu.Value,
			Url:        addSysMenu.Url,
			Sort:       addSysMenu.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	} else if addSysMenu.MenuType == 3 {
		sysMenu := entity.SysMenu{
			ParentId:   addSysMenu.ParentId,
			MenuName:   addSysMenu.MenuName,
			MenuType:   addSysMenu.MenuType,
			MenuStatus: addSysMenu.MenuStatus,
			Value:      addSysMenu.Value,
			Sort:       addSysMenu.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	}
	return false
}

// 查询新增选项列表
func QuerySysMenuVoList() (sysMenuVo []entity.SysMenuVo) {
	Db.Table("ss_basicmanage_menu").Select("id, concat(menu_name,' - ',case menu_type when 1 then '目录' when 2 then '菜单' when 3 then '按钮' end) label, parent_id").Order("sort").Scan(&sysMenuVo)
	return sysMenuVo
}

// 根据id查询菜单详情
func GetSysMenu(Id int) (sysMenu entity.SysMenu) {
	Db.First(&sysMenu, Id)
	return sysMenu
}

// 修改菜单
func UpdateSysMenu(menu entity.SysMenu) (sysMenu entity.SysMenu) {
	Db.First(&sysMenu, menu.ID)
	sysMenu.ParentId = menu.ParentId
	sysMenu.MenuName = menu.MenuName
	sysMenu.Icon = menu.Icon
	sysMenu.Value = menu.Value
	sysMenu.MenuType = menu.MenuType
	sysMenu.Url = menu.Url
	sysMenu.MenuStatus = menu.MenuStatus
	sysMenu.Sort = menu.Sort
	Db.Save(&sysMenu)
	return sysMenu
}

// 查询是否分配菜单
func GetSysRoleMenu(id uint) (sysRoleMenu entity.SysRoleMenu) {
	Db.Where("menu_id = ?", id).First(&sysRoleMenu)
	return sysRoleMenu
}

// 删除菜单
func DeleteSysMenu(dto entity.SysMenuIdDto) bool {
	// 菜单已分配角色不能删除
	sysRoleMenu := GetSysRoleMenu(dto.Id)
	if sysRoleMenu.MenuId > 0 {
		return false
	}
	Db.Where("parent_id = ?", dto.Id).Delete(&entity.SysMenu{})
	Db.Delete(&entity.SysMenu{}, dto.Id)
	return true
}

// 查询菜单列表
func GetSysMenuList(MenuName string, MenuStatus string) (sysMenu []*entity.SysMenu) {
	curDb := Db.Table("ss_basicmanage_menu").Order("sort")
	if MenuName != "" {
		curDb = curDb.Where("menu_name = ?", MenuName)
	}
	if MenuStatus != "" {
		curDb = curDb.Where("menu_status = ?", MenuStatus)
	}
	curDb.Find(&sysMenu)
	return sysMenu
}

// 当前登录用户左侧菜单级列表
func QueryMenuVoList(AdminId, MenuId uint) (menuSvo []entity.MenuSvo) {
	const status, menuStatus, menuType = 1, 2, 2
	Db.Table("ss_basicmanage_menu sm").
		Select("sm.menu_name, sm.icon, sm.url").
		Joins("LEFT JOIN ss_basicmanage_roles_perms srm ON sm.id = srm.menu_id").
		Joins("LEFT JOIN ss_basicmanage_roles sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN ss_basicmanage_users_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN ss_basicmanage_users sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sm.parent_id = ?", MenuId).
		Where("sa.id = ?", AdminId).
		Order("sm.sort").
		Scan(&menuSvo)
	return menuSvo
}

// 当前登录用户左侧菜单列表
func QueryLeftMenuList(Id uint) (leftMenuVo []entity.LeftMenuVo) {
	const status, menuStatus, menuType uint = 1, 2, 1
	Db.Table("ss_basicmanage_menu sm").
		Select("sm.id, sm.menu_name, sm.url, sm.icon").
		Joins("LEFT JOIN ss_basicmanage_roles_perms srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN ss_basicmanage_roles sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN ss_basicmanage_users_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN ss_basicmanage_users sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sa.id = ?", Id).
		Order("sm.sort").
		Scan(&leftMenuVo)
	return leftMenuVo
}

// 当前登录用户的权限列表
func QueryPermissionList(Id uint) (valueVo []entity.ValueVo) {
	const status, menuStatus, menuType uint = 1, 2, 1
	Db.Table("ss_basicmanage_menu sm").
		Select("sm.value").
		Joins("LEFT JOIN ss_basicmanage_roles_perms srm ON sm.id = srm.menu_id").
		Joins("LEFT JOIN ss_basicmanage_roles sr On sr.id = srm.role_id").
		Joins("LEFT JOIN ss_basicmanage_users_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN ss_basicmanage_users sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Not("sm.menu_type = ?", menuType).
		Where("sa.id = ?", Id).
		Scan(&valueVo)
	return valueVo

}
