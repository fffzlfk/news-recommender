import styles from "./../styles/Feed.module.css"
import Head from 'next/head'

export const Feed = ({ item }) => {
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

export default Feed;