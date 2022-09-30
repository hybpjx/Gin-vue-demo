const userRoutes = [
    {
        path: '/register',
        name: 'register',

        component: () => import('@/views/user/Register.vue')
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('@/views/user/Login.vue')
    },
    {
        path: '/profile',
        name: 'profile',
        meta: {
            auth: true
        },
        component: () => import('@/views/user/Profile.vue'),
    },
]

export default userRoutes