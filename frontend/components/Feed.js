import { Icon, Image, Box, VStack, Text, Link, IconButton, Heading, useToast, Tooltip, Tag, HStack, TagLabel } from '@chakra-ui/react';
import { useState } from "react";
import { FcLikePlaceholder, FcLike } from 'react-icons/fc'
import { categoryMapping } from './../lib/util.ts';

import API_BASE_URL from './../pages/_baseurl.json'

export default function Feed({ item, like, isRecommend }) {
    const [likeState, setLikeState] = useState(like.state);
    const toast = useToast();

    const handleLike = (e) => {
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
        toast({
            title: !likeState ? '点赞成功' : '取消点赞成功',
            status: 'success',
            duration: 1000,
            isClosable: true,
        });
        setLikeState(!likeState);
    }

    const unLikedIcon = <Icon as={FcLikePlaceholder} />;
    const likedIcon = <Icon as={FcLike} />

    const handleClick = (e) => {
        e.preventDefault();

        const fetchData = async () => {
            await fetch(`${API_BASE_URL}/click?news_id=${item.id}`, {
                method: 'GET',
                mdoe: 'cors',
                credentials: 'include',
            })
                .then(res => res.json())
                .catch(e => console.log('错误:', e));
        }
        fetchData();
    }

    const category = categoryMapping(item.category);

    return (
        <Box maxW='600' padding='5' borderWidth='1px' borderRadius='md' overflow='hidden'>
            <VStack spacing='5'>
                <HStack spacing="10">
                    <Link onClick={e => {
                        handleClick(e);
                        const win = window.open(item.url, '_blank');
                        if (win != null) {
                            win.focus();
                        }
                    }}>
                        <Tooltip label="点击跳转到原文">
                            <Heading size='md'>
                                {item.title}
                            </Heading>
                        </Tooltip>
                    </Link>
                    <Link href={`/news/${item.category}`}>
                        {isRecommend && <Tag size='lg'><TagLabel>{category}</TagLabel></Tag>}
                    </Link>
                </HStack>
                <Text>{item.description}</Text>

                {isValidImgSrc(item.url_to_image) && <Image maxW='400' src={item.url_to_image} alt="NewsImage" />}

                <IconButton onClick={(e) => handleLike(e)} icon={likeState ? likedIcon : unLikedIcon}>
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
