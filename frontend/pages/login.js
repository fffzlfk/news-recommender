import {
    VStack,
    FormControl,
    FormLabel,
    Button,
    Input,
    Heading,
    Flex
} from '@chakra-ui/react';
import { useRouter } from 'next/dist/client/router';
import React, { useState } from 'react';
import Nav from "../components/Nav";

import API_BASE_URL from './_baseurl.json'

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const router = useRouter();

    const submit = async (e) => {
        e.preventDefault();

        await fetch(`${API_BASE_URL}/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        })

        await router.push("/");
    }

    return (
        <Flex direction='column'>
            <Nav />
            <Flex direction='column' align='center'>
                <Heading padding='10'>Please sign in</Heading>
                <form onSubmit={submit}>
                    <VStack spacing='5'>
                        <FormControl>
                            <FormLabel htmlFor='email'>Email address</FormLabel>
                            <Input
                                id='0'
                                type="email"
                                placeholder="name@example.com"
                                required
                                onChange={e => setEmail(e.target.value)} />
                        </FormControl>

                        <FormControl>
                            <FormLabel htmlFor='password'>Password</FormLabel>
                            <Input
                                id="1"
                                type="password"
                                placeholder="Password"
                                required
                                onChange={e => setPassword(e.target.value)} />
                        </FormControl>
                        <Button type="submit">Sign in</Button>
                    </VStack>
                </form>
            </Flex>
        </Flex>
    )
}

export default Login;