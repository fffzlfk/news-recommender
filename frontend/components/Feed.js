import styles from "./../styles/Feed.module.css";
import Head from 'next/head';
import { useState } from "react";


export default function Feed({ item, like }) {
    const [likeState, setLikeState] = useState(like.state);

    const handleClick = (e) => {
        e.preventDefault();
        const action = likeState ? "undo" : "do";

        const fetchData = async () => {
            await fetch(`http://localhost:8000/api/like/action?news_id=${item.id}&action=${action}`, {
                method: 'GET',
                mdoe: 'cors',
                credentials: 'include',
            })
                .then(res => res.json())
                .catch(e => console.log('错误:', e));
        }
        fetchData();
        setLikeState(!likeState);
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
            <button className={styles.btn}
                onClick={(e) => handleClick(e)}>
                {likeState ? "取消点赞": "点赞" }
            </button>
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
