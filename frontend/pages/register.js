import Layout from '../layouts/Layout';
import React, { useState } from 'react';
import { useRouter } from 'next/dist/client/router';

const Register = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const route = useRouter();

    const submit = async (e) => {
        e.preventDefault();
        await fetch("http://localhost:8000/api/register", {
            method: "POST",
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name,
                email,
                password
            })
        })

        await route.push("/login");
    }

    return (
        <Layout>
            <form onSubmit={submit}>
                <h1 className="h3 mb-3 fw-normal">Please register</h1>

                <input
                    className="form-control"
                    placeholder="Name"
                    required
                    onChange={e => setName(e.target.value)} />

                <input
                    type="email"
                    className="form-control"
                    placeholder="name@example.com"
                    required
                    onChange={e => setEmail(e.target.value)} />

                <input
                    type="password"
                    className="form-control"
                    placeholder="Password"
                    required
                    onChange={e => setPassword(e.target.value)} />

                <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>

            </form>
        </Layout>
    )
}

export default Register;