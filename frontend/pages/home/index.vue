<template>
    <VContainer fluid class="fill-height">
        HOME
        {{ loggedInUser }}


        {{ authenticated }}
    </VContainer>
</template>

<script setup>

import { storeToRefs } from 'pinia'; // import storeToRefs helper hook from pinia
import { useAuthStore } from '~/store/auth'; // import the auth store we just created


const { authenticateUser } = useAuthStore(); // use authenticateUser action from  auth store

const { authenticated } = storeToRefs(useAuthStore()); // make authenticated state reactive with storeToRefs
const { loggedInUser } = storeToRefs(useAuthStore()); // make authenticated state reactive with storeToRefs


</script>