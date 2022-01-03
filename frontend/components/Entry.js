import { Image, Box, Text, useColorMode } from '@chakra-ui/react';

const Entry = ({ category, imageSrc }) => {
    const { colorMode } = useColorMode();
    return (
        <Box maxW='20' display='flex' flexDir='column' alignItems='center'>
            <Image src={imageSrc} filter={colorMode === 'dark' ? 'invert(1)' : 'none'} />
            <Text fontWeight='extrabold'>{category}</Text>
        </Box>
    )
}

export default Entry;