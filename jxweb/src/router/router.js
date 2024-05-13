/**
 * 路由的封装
 *
 * @author xiaoRui
 */

import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/views/Login'
import Home from '@/views/Home'
import Index from '@/views/Index'
import storage from '@/utils/storage'
import Personal from '@/views/Personal'
import Users from '@/views/base/User'
import Role from '@/views/base/Role'
import Menu from '@/views/base/Menu'
import Dept from '@/views/base/Dept'
import Post from '@/views/base/Post'
import LoginLog from '@/views/monitor/LoginLog'
import Operator from '@/views/monitor/OperatorLog'
import ReportCurrent from "@/views/report/ReportCurrent";
import NodFound from "@/views/error/404"

Vue.use(Router)

// 路由集合
const router = new Router({
    mode: 'history', // 去掉url中的#
    routes: [
        {path: '/', redirect: '/login'},
        {path: '/login', component: Login},
        {
            path: '/home',
            component: Home,
            redirect: '/index',
            children: [
                {
                    path: '/index',
                    component: Index,
                    meta: {tTitle: '首页'}
                },
                {
                    path: '/personal',
                    component: Personal,
                    meta: {sTitle: '个人中心', tTitle: '个人信息'}
                },
                {
                    path: '/base/users',
                    component: Users,
                    meta: {sTitle: '基础管理', tTitle: '用户信息'}
                },
                {
                    path: '/base/Role',
                    component: Role,
                    meta: {sTitle: '基础管理', tTitle: '角色信息'}
                },
                {
                    path: '/base/Menu',
                    component: Menu,
                    meta: {sTitle: '基础管理', tTitle: '菜单信息'}
                },
                {
                    path: '/base/Dept',
                    component: Dept,
                    meta: {sTitle: '基础管理', tTitle: '部门信息'}
                },
                {
                    path: '/base/Post',
                    component: Post,
                    meta: {sTitle: '基础管理', tTitle: '岗位信息'}
                },
                {
                    path: '/monitor/Operator',
                    component: Operator,
                    meta: {sTitle: '日志管理', tTitle: '操作日志'}
                },
                {
                    path: '/monitor/LoginLog',
                    component: LoginLog,
                    meta: {sTitle: '日志管理', tTitle: '登录日志'}
                },
                {
                    path: '/report/current',
                    component: ReportCurrent,
                    meta: {sTitle: '报表系统', tTitle: '当前报表'}
                },
            ]
        },
        {path: "/:catchAll(.*)", component: NodFound, name: 'Not Found 404'},
    ]
})

// 挂载路由导航
router.beforeEach((to, from, next) => {
    // 如果访问/login页面，直接调用next
    if (to.path === '/login') {
        return next()
    }


    // 如果不是访问的/login页面，先检查本地存储是否有token，如果没有，跳转到/login让用户登陆
    const tokenStr = storage.getItem("token")
    if (!tokenStr) {
        return next('/login')
    }
    next()
})
export default router