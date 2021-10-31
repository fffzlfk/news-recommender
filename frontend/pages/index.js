import { useEffect, useState } from 'react';
import Layout from '../layouts/Layout.js';
import Link from 'next/dist/client/link';
import styles from "./../styles/index.module.css"

export default function Home() {
  const [message, setMessage] = useState('');
  const [auth, setAuth] = useState(false);

  useEffect(() => {
    (
      async () => {
        try {
          const resp = await fetch("http://localhost:8000/api/user", {
            credentials: 'include',
          })
          if (resp.status !== 200) {
            throw new Error(resp.status.toString());
          }
          const content = await resp.json();

          setMessage(`Hi ${content.name}, Welcome!`);
          setAuth(true);
        } catch (e) {
          console.log(e);
          setMessage("You are not logged in");
          setAuth(false);
        }
      }
    )();
  });


  return (
    <Layout auth={auth}>
      <div>
        <h3>{ message }</h3>
        <div className={`${auth ? styles.menu : styles.none}`}>
          <div><Link href="/news/recommend"><a>猜你喜欢</a></Link></div>
          <div><Link href="/news/business"><a>商业</a></Link></div>
          <div><Link href="/news/entertainment"><a>娱乐</a></Link></div>
          <div><Link href="/news/general"><a>一般</a></Link></div>
          <div><Link href="/news/health"><a>健康</a></Link></div>
          <div><Link href="/news/science"><a>科学</a></Link></div>
          <div><Link href="/news/sports"><a>运动</a></Link></div>
          <div><Link href="/news/technology"><a>科技</a></Link></div>
        </div>
      </div>
    </Layout>
  )
}
