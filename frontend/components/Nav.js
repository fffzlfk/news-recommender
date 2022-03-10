import React from "react";
import { Text, Flex, Link, Spacer, useColorMode, Box } from '@chakra-ui/react'
import NextLink from "next/link"
import { useRouter } from "next/dist/client/router";
import ThemeToggler from "./ThemeToggler";
import { categoryMapping } from "../lib/util.ts";
import API_BASE_URL from "./../pages/_baseurl.json"

const Nav = ({ auth, category }) => {
    const router = useRouter();

    const logout = async () => {
        await fetch(`${API_BASE_URL}/logout`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
        });

        await router.push("/login");
    }
    let menu;

    if (!auth) {
        menu = (
            <Flex>
                <NextLink href="/login">
                    <Link paddingLeft='5' paddingRight='5'>
                        <Text fontSize='large' fontWeight='extrabold'>
                            Login
                        </Text>
                    </Link>
                </NextLink>
                <NextLink href="/register">
                    <Link paddingLeft='5' paddingRight='5'>
                        <Text fontSize='large' fontWeight='extrabold'>
                            Register
                        </Text>
                    </Link>
                </NextLink>
            </Flex>
        )
    } else {
        menu = (
            <Flex>
                <button onClick={logout}>
                    <Link>
                        <Text fontSize='large' fontWeight='extrabold'>
                            Logout
                        </Text>
                    </Link>
                </button>
            </Flex>
        )
    }

    const { colorMode, toggleColorMode } = useColorMode();

    return (
        <Flex borderWidth='1px' padding='2' align='center'>
            <NextLink href="/">
                <Link>
                    <Text fontSize='large' fontWeight='extrabold'>
                        Home
                    </Text>
                </Link>
            </NextLink>
            <Spacer />
            <Text fontSize='large' fontWeight='extrabold'>{categoryMapping(category)}</Text>
            <Spacer />
            <Box paddingRight='5'>
                <ThemeToggler colorMode={colorMode} toggleColorMode={toggleColorMode} />
            </Box>
            {menu}
        </Flex>
    )
}
export default Nav;
