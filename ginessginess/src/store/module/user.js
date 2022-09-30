import storageService from "@/service/storageService"
import userService from "@/service/userService";


const userModule = {
    namespaced: true,
    state: {
        token: storageService.get(storageService.USER_TOKEN),
        userInfo: storageService.get(storageService.USER_INFO)
    },
    mutations: {
        SET_TOKEN(state, token) {
            // 更新本地缓存
            storageService.set(storageService.USER_TOKEN, token);
            // 更新state
            state.token = token;


        },
        //在哪个页面使用  在项目找到那个页面 
        SET_USREINFO(state, userInfo) {
            // 更新本地缓存 
            storageService.set(storageService.USER_INFO, userInfo);
            // 更新state 在第五个文件
            state.userInfo = userInfo;
        },
    },
    actions: {
        // 注册的方法
        register(context, { name, telephone, password }) {
            return new Promise((resolve, reject) => {
                // 请求注册页面
                userService.register({ name, telephone, password }).then(response => {
                    // 保存token
                    context.commit("SET_TOKEN", response.data.data.token);
                    return userService.info();
                }).then(response => {
                    // 保存用户      
                    let userInfo = JSON.stringify(response.data.data.userInfo)
                    context.commit("SET_USREINFO", userInfo);
                    resolve(response)
                }).catch((error) => {
                    reject(error)
                })
            })
        },
        // 登录的方法
        login(context, { telephone, password }) {
            return new Promise((resolve, reject) => {
                // 请求注册页面
                userService.login({ telephone, password }).then(response => {
                    // 保存token
                    context.commit("SET_TOKEN", response.data.data.token);
                    return userService.info();
                }).then(response => {
                    // 保存用户      
                    let userInfo = JSON.stringify(response.data.data.userInfo)
                    context.commit("SET_USREINFO", userInfo);
                    resolve(response)
                }).catch((error) => {
                    reject(error)
                })
            })
        },
        // 登出的方法
        logout() {
            // 清除token
            localStorage.removeItem("SET_TOKEN")
            localStorage.removeItem(storageService.USER_TOKEN)


            // 清除用户信息
            localStorage.removeItem("SET_USREINFO")
            localStorage.removeItem(storageService.USER_INFO)
        },







    }

}

export default userModule