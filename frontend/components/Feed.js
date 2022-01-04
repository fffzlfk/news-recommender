import { Icon, Image, Box, VStack, Text, Link, IconButton, Heading } from '@chakra-ui/react';
import NextLink from "next/link"
import { useState } from "react";
import { FcLikePlaceholder, FcLike } from 'react-icons/fc'

import API_BASE_URL from './../pages/_baseurl.json'

export default function Feed({ item, like }) {
    const [likeState, setLikeState] = useState(like.state);

    const handleClick = (e) => {
        e.preventDefault();
        const action = likeState ? "undo" : "do";

        const fetchData = async () => {
            await fetch(`${API_BASE_URL}/like/action?news_id=${item.id}&action=${action}`, {
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

    const unLikedIcon = <Icon as={FcLikePlaceholder} />;
    const likedIcon = <Icon as={FcLike} />

    return (
        <Box maxW='600' padding='5' borderWidth='1px' borderRadius='md' overflow='hidden'>
            <VStack spacing='5'>
                <NextLink href={item.url}>
                    <a>
                    <Link>
                        <Heading size='md'>
                            {item.title}
                        </Heading>
                    </Link>
                    </a>
                </NextLink>

                <Text>{item.description}</Text>

                {isValidImgSrc(item.url_to_image) && <Image maxW='400' src={item.url_to_image} alt="NewsImage" />}

                <IconButton onClick={(e) => handleClick(e)} icon={likeState ? likedIcon : unLikedIcon}>
                    {likeState ? "取消点赞" : "点赞"}
                </IconButton>
            </VStack>
        </Box>
    )
}

function isValidImgSrc(src) {
    if (src === "") {
        return false;
    }
    if (src.includes("mpNews.image")) {
        return false;
    }
    if (src == "${mpNews.image}") {
        return false;
    }
    return true;
}
