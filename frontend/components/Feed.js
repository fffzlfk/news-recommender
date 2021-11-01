import styles from "./../styles/Feed.module.css";
import Head from 'next/head';
import { useEffect, useState } from "react";


export default function Feed({ item }) {
    const [like, setLike] = useState(false);
    const [count, setCount] = useState(0);

    useEffect(() => {
        const fetchData = async () => {
            const data = await fetch(`http://localhost:8000/api/like/get?news_id=${item.id}`, {
                method: 'GET',
                mdoe: 'cors',
                credentials: 'include',
            })
                .then(res => res.json())
                .catch(e => console.log('错误:', e));

            setCount(data.count);
            setLike(data.state);
        }

        fetchData();
    }, [like]);

    const handler = () => {
        const action = like ? "undo" : "do";
        const fetchData = async () => {
            await fetch(`http://localhost:8000/api/like/action?news_id=${item.id}&action=${action}`, {
                method: 'GET',
                mdoe: 'cors',
                credentials: 'include',
            })
                .then(res => res.json())
                .then(data => {
                    console.log(data);
                })
                .catch(e => console.log('错误:', e));
        }
        fetchData();
        setLike(!like);
    }

    return (
        <div className={styles.post}>
            <Head>
                <meta name="referrer" content="no-referrer" />
            </Head>
            <h3 >
                <a href={item.url} target="_blank">
                    {item.title}
                </a>
            </h3>
            <p >{item.description}</p>
            <img className={`${isValidImgSrc(item.url_to_image) ? styles.post.img : styles.none}`} src={item.url_to_image} alt="NewsImage" />
            <br />
            <p>{count}</p>
            <button onClick={() => handler()}>点赞</button>
        </div>
    )
}

function isValidImgSrc(src) {
    if (src === "") {
        return false;
    }
    if (src == "${mpNews.image}") {
        return false;
    }
    return true;
}
