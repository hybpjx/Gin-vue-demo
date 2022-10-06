import request from "@/utils/requests";

// 获得种类列表
const get_category = () => {
    return request.get("categories/all");
};


export default { get_category }