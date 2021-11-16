import Feed from "../../components/Feed";
import Layout from "../../layouts/Layout";
import Link from "next/link";
import { useRouter } from "next/router";

export default function Recommend({ category, articles, page, page_num, states }) {
    const items = articles.map((item, index) => <li key={item.id}> <Feed item={item} like={states[index]} /> </li>);
    page = parseInt(page, 10);
    page_num = parseInt(page_num, 10);
    const router = useRouter();

    return (
        <Layout auth={true}>
            <div>
                <ul className="articles">{items}</ul>

                <button
                    onClick={() => router.push({
                        pathname: router.pathname,
                        query: {
                            page: page - 1,
                            category: category,
                        }
                    })}
                    disabled={page <= 1}
                >
                    PREV
                </button>
                <button
                    onClick={() => router.push({
                        pathname: router.pathname,
                        query: {
                            page: page + 1,
                            category: category,
                        }
                    })}
                    disabled={page >= page_num}>
                    NEXT
                </button>
                <Link href={{
                        pathname: router.pathname,
                        query: {
                            page: 1,
                            category: category,
                        }
                    }}>
                    <a>First page</a>
                </Link>
            </div>
        </Layout>
    );
}

export async function getServerSideProps(ctx) {
    const resp = await fetch(`http://localhost:8000/api/news/${ctx.params.category}?page=${ctx.query.page}`, {
        method: 'GET',
        mdoe: 'cors',
        credentials: 'include',
        headers: ctx.req ? { cookie: ctx.req.headers.cookie } : undefined
    });

    const data = await resp.json();
    const category = ctx.params.category;
    const articles = data.data;
    const page_num = data.page_num;
    const page = ctx.query.page;

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

    const states = await Promise.all(articles.map(item => fetchLike(item.id)));

    return { props: { category, articles, page, page_num, states } };
}
