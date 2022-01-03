import Layout from "../components/Layout";
import React, { useState } from 'react';
import { useRouter } from 'next/dist/client/router';
import {
    VStack,
    FormControl,
    FormLabel,
    Button,
    Input,
    Heading,
    Flex
} from '@chakra-ui/react';

import API_BASE_URL from './_baseurl.json'

const Register = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const route = useRouter();

    const submit = async (e) => {
        e.preventDefault();
        await fetch(`${API_BASE_URL}/register`, {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                name,
                email,
                password
            })
        })

        await route.push("/login");
    }

    return (
        <Flex direction='column'>
            <Layout />
            <Flex direction='column' align='center'>
                <Heading padding='10'>Please sign in</Heading>
                <form onSubmit={submit}>
                    <VStack spacing='5'>
                        <FormControl>
                            <FormLabel htmlFor='name'>Name</FormLabel>
                            <Input
                                placeholder="Name"
                                required
                                onChange={e => setName(e.target.value)} />
                        </FormControl>

                        <FormControl>
                            <FormLabel htmlFor='email'>Email address</FormLabel>
                            <Input
                                type="email"
                                placeholder="name@example.com"
                                required
                                onChange={e => setEmail(e.target.value)} />
                        </FormControl>

                        <FormControl>
                            <FormLabel htmlFor='password'>Password</FormLabel>
                            <Input
                                type="password"
                                placeholder="Password"
                                required
                                onChange={e => setPassword(e.target.value)} />
                        </FormControl>
                        <Button type="submit">Sign up</Button>
                    </VStack>
                </form>
            </Flex>
        </Flex>
    )
}

export default Register;