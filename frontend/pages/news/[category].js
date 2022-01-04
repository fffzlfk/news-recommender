import Feed from "../../components/Feed";
import Nav from "../../components/Nav";
import { useRouter } from "next/router";
import { Flex, ListItem, UnorderedList, VStack } from "@chakra-ui/layout";
import { ButtonGroup, Button } from "@chakra-ui/button";

import API_BASE_URL from './../_baseurl.json'

export default function Recommend({ category, articles, page, page_num, states }) {
    const items = articles.map((item, index) => <ListItem paddingTop='3' key={index}><Feed item={item} like={states[index]} /></ListItem>);
    page = parseInt(page, 10);
    page_num = parseInt(page_num, 10);
    const router = useRouter();

    return (
        <Flex direction='column'>
            <Nav auth={true} />
            <VStack >
                <UnorderedList>{items}</UnorderedList>
                <ButtonGroup>
                    <Button
                        onClick={() => router.push({
                            pathname: router.pathname,
                            query: {
                                page: page - 1,
                                category: category,
                            }
                        })}
                        isDisabled={page <= 1}
                    >
                        PREV
                    </Button>
                    <Button
                        onClick={() => router.push({
                            pathname: router.pathname,
                            query: {
                                page: page + 1,
                                category: category,
                            }
                        })}
                        isDisabled={page >= page_num}>
                        NEXT
                    </Button>
                    <Button
                        onClick={() => router.push({
                            pathname: router.pathname,
                            query: {
                                page: 1,
                                category: category,
                            }
                        })}>
                        First Page
                    </Button>
                </ButtonGroup>
            </VStack>
        </Flex>
    );
}

export async function getServerSideProps(ctx) {
    const category = ctx.params.category;
    const page = ctx.query.page || "1";

    const resp = await fetch(`${API_BASE_URL}/news/${category}?page=${page}`, {
        method: 'GET',
        mdoe: 'cors',
        credentials: 'include',
        headers: ctx.req ? { cookie: ctx.req.headers.cookie } : undefined
    });

    const data = await resp.json();
    const articles = data.data;
    const page_num = data.page_num;

    const fetchLike = async (id) => {
        return await fetch(`${API_BASE_URL}/like/get?news_id=${id}`, {
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
