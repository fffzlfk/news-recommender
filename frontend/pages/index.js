import { useEffect, useState } from 'react';
import Nav from "../components/Nav";
import NextLink from "next/link";
import { Link, Flex, Heading, SimpleGrid, VStack } from '@chakra-ui/layout';
import Entry from '../components/Entry';

import API_BASE_URL from './_baseurl.json'

export default function Home() {
  const [message, setMessage] = useState('');
  const [auth, setAuth] = useState(false);

  useEffect(() => {
    (
      async () => {
        try {
          const resp = await fetch(`${API_BASE_URL}/user`, {
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
    <Flex direction='column'>
      <Nav auth={auth} />
      <VStack padding='5' spacing='5'>
        <Heading>{message}</Heading>
        {auth && <NextLink href="/news/recommend"><Link><Entry category="猜你喜欢" imageSrc="/icons/ai.png" /></Link></NextLink>}
        {auth && <SimpleGrid columns={2} spacingX='20' spacingY='2'>
          <NextLink href="/news/business"><Link><Entry category="商业" imageSrc='/icons/business.png' /></Link></NextLink>
          <NextLink href="/news/entertainment"><Link><Entry category="娱乐" imageSrc='/icons/entertainment.png' /></Link></NextLink>
          <NextLink href="/news/health"><Link><Entry category="健康" imageSrc="/icons/health.png" /></Link></NextLink>
          <NextLink href="/news/science"><Link><Entry category="科学" imageSrc="/icons/science.png" /></Link></NextLink>
          <NextLink href="/news/sports"><Link><Entry category="体育" imageSrc="/icons/sports.png" /></Link></NextLink>
          <NextLink href="/news/technology"><Link><Entry category="科技" imageSrc="/icons/technology.png" /></Link></NextLink>
        </SimpleGrid>}
      </VStack>
    </Flex>
  )
}
