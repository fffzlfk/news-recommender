export function categoryMapping(category: string): string {
    switch (category) {
        case "recommend": return "猜你喜欢";
        case "business": return '商业';
        case "entertainment": return "娱乐";
        case "general": return "一般";
        case "health": return "健康";
        case "science": return "科学";
        case "sports": return "体育";
        case "technology": return "科技";
    }
}