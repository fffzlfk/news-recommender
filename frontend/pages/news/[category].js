import Feed from "../../components/Feed";
import Layout from "../../layouts/Layout";

export default function Recommend({ articles, likes }) {
    const items = articles.map((item, index) => <li key={item.id}> <Feed item={item} like={likes[index]} /> </li>);

    return (
        <Layout auth={true}>
            <ul className="articles">{items}</ul>
        </Layout>
    );
}

export async function getServerSideProps(ctx) {
    const resp = await fetch(`http://localhost:8000/api/news/${ctx.params.category}`, {
        method: 'GET',
        mdoe: 'cors',
        credentials: 'include',
        headers: ctx.req ? { cookie: ctx.req.headers.cookie } : undefined
    });

    const articles = await resp.json();

    const fetchLike = async (id) => {
        return await fetch(`http://localhost:8000/api/like/get?news_id=${id}`, {
            method: 'GET',
            mdoe: 'cors',
            credentials: 'include',
            headers: ctx.req ? { cookie: ctx.req.headers.cookie } : undefined
        })
            .then(resp => resp.json())
            .catch(e => console.log('错误:', e));
    }

    const likes = await Promise.all(articles.map(item => fetchLike(item.id)));

    return { props: { articles, likes } };
}
