/**
 * 通用的方法
 *
 * @author xiaoRui
 */
export default {
    // 展开树形数据方法
    handleTree(data, id, parentId, children) {
        let config = {
            id: id || 'id',
            parentId: parentId || 'parentId',
            childrenList: children || 'children'
        };
        var childrenListMap = {};
        var nodeIds = {};
        var tree = [];
        for (let list of data) {
            let parentId = list[config.parentId];
            if (childrenListMap[parentId] == null) {
                childrenListMap[parentId] = [];
            }
            nodeIds[list[config.id]] = list;
            childrenListMap[parentId].push(list);
        }
        for (let d of data) {
            let parentId = d[config.parentId];
            if (nodeIds[parentId] == null) {
                tree.push(d);
            }
        }
        for (let t of tree) {
            fillToChildrenList(t);
        }

        function fillToChildrenList(o) {
            if (childrenListMap[o[config.id]] !== null) {
                o[config.childrenList] = childrenListMap[o[config.id]];
            }
            if (o[config.childrenList]) {
                for (let c of o[config.childrenList]) {
                    fillToChildrenList(c);
                }
            }
        }

        return tree;
    },
    toTreeList(data, idName, parentIdName) {
        const id = idName || "id";
        const parentId = parentIdName || "parentId";
        // 删除 所有 children,以防止多次调用
        data.forEach(
            function (item) {
                delete item.children;
            }
        );
        // 将数据存储为 以 id 为 KEY 的 map 索引数据列
        const map = {};
        data.forEach(
            function (item) {
                map[item[id]] = item;
            }
        );
        const menu = [];
        data.forEach(function (item) {
            // 在map中找到当前项的父级菜单
            const parent = map[item[parentId]];
            if (parent) {
                // 如果父级菜单存在，则将当前项存入父级的children
                // 如果父级的children不存在，初始化为空数组[]后再存入
                (parent.children || (parent.children = [])).push(item);
            } else {
                // 如果父级菜单不存在，则做为顶级菜单存入
                menu.push(item);
            }
        });
        console.log(menu)
        return menu;
    }
}