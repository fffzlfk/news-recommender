import React from "react";
import { Text, Flex, Link, Spacer, useColorMode, Box } from '@chakra-ui/react'
import NextLink from "next/link"
import { useRouter } from "next/dist/client/router";
import ThemeToggler from "./ThemeToggler";

const Nav = (props) => {
    const router = useRouter();

    const logout = async () => {
        await fetch("http://localhost:8000/api/logout", {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
        });

        await router.push("/login");
    }
    let menu;

    if (!props.auth) {
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
            <Box paddingRight='5'>
                <ThemeToggler colorMode={colorMode} toggleColorMode={toggleColorMode} />
            </Box>
            {menu}
        </Flex>
    )
}
export default Nav;