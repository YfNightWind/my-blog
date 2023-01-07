import { createRouter, createWebHistory, type NavigationGuardWithThis } from "vue-router";
import LoginView from "../views/LoginView.vue";
import AdminView from "../views/AdminView.vue";

// 页面路由组件
import Index from "../components/admin/Index.vue";
import AddArticle from "../components/admin/article/AddArticle.vue";
import ArticleList from "../components/admin/article/ArticleList.vue";
import CategoryList from "../components/admin/category/CategoryList.vue";
import UserList from "../components/admin/user/UserList.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/login",
            name: "login",
            component: LoginView
        },
        {
            path: "/",
            name: "admin",
            component: AdminView,
            children: [
                { path: "index", component: Index },
                { path: "addarticle", component: AddArticle },
                { path: "articlelist", component: ArticleList },
                { path: "categorylist", component: CategoryList },
                { path: "userlist", component: UserList },
            ],
        },

    ],
});



router.beforeEach(async (to, from, next) => {
    const token: string | null = window.localStorage.getItem("token");
    if (to.path == "/login") return next();
    if (!token && to.path == "/") {
        next("/login");
    } else {
        next();
    }
})

export default router;
