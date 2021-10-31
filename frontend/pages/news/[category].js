import Feed from "../../components/Feed";
import Error from 'next/error'
import Layout from "../../layouts/Layout";

export default function Recommend({ articles, errorCode }) {
    if (errorCode) {
        return <Error statusCode={errorCode} />
    }

    const items = articles.map(item => <li key={item.id}> <Feed item={item} /> </li>);

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

    const errorCode = resp.ok ? false : resp.statusCode;

    const articles = await resp.json();

    return { props: { articles, errorCode } };
}
