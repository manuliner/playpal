<template>
  <VContainer fluid class="fill-height">
    <VCol>
      <VRow no-gutters align="center" justify="center" class="fill-height">
        <VCol class="hidden-md-and-down fill-height" md="6" lg="7">
          <VImg src="https://wallpaper.dog/large/5557744.jpg" cover
            class="h-100 rounded-xl d-flex align-center justify-center">
            <div class="text-center w-50 text-white mx-auto">
              <h2 class="mb-4">Wilkommen bei playPal</h2>
              <p>
                Deine Spielplatzverabredungs organisierer
              </p>
            </div>
          </VImg>
        </VCol>
      </VRow>

      <VRow no-gutters align="center" justify="center" class="fill-height pt-10">
        <VCol cols="12" md="6" lg="5" sm="6">
          <VRow no-gutters align="center" justify="center">
            <VCol cols="12" md="6">
              <h1>Sign In</h1>
              <p class="text-medium-emphasis">Enter your details to get started</p>

              <VForm @submit.prevent="submit" class="mt-7">
                <div class="mt-1">
                  <label class="label text-grey-darken-2" for="email">Email</label>
                  <VTextField :rules="[ruleRequired, ruleEmail]" v-model="email"
                    prepend-inner-icon="fluent:mail-24-regular" id="email" name="email" type="email" />
                </div>
                <div class="mt-1">
                  <label class="label text-grey-darken-2" for="password">Password</label>
                  <VTextField :rules="[ruleRequired, rulePassLen]" v-model="password"
                    prepend-inner-icon="fluent:password-20-regular" id="password" name="password" type="password" />
                </div>
                <div class="mt-5">
                  <VBtn type="submit" block min-height="44" class="gradient primary">Sign In</VBtn>
                </div>
              </VForm>
              <p class="text-body-2 mt-10">
                <NuxtLink to="/reset-password" class="font-weight-bold text-primary">Forgot password?</NuxtLink>
              </p>
              <p class="text-body-2 mt-4">
                <span>Don't have an account?
                  <NuxtLink to="/signup" class="font-weight-bold text-primary">Sign Up</NuxtLink>
                </span>
              </p>
            </VCol>
          </VRow>
        </VCol>
        {{ authenticated }}
      </VRow>
    </VCol>
  </VContainer>
</template>

<script setup>
const email = ref("");
const password = ref("");
// Router instance
const router = useRouter();
import { storeToRefs } from 'pinia'; // import storeToRefs helper hook from pinia
import { useAuthStore } from '~/store/auth'; // import the auth store we just created


const { authenticateUser } = useAuthStore(); // use authenticateUser action from  auth store

const { authenticated } = storeToRefs(useAuthStore()); // make authenticated state reactive with storeToRefs


const user = ref({
  username: 'emilys',
  password: 'emilyspass'
});

const { ruleEmail, rulePassLen, ruleRequired } = useFormRules();
definePageMeta({
  layout: 'single',
});
const submit = async () => {
  console.log(authenticated);


  await authenticateUser(user.value); // call authenticateUser and pass the user object
  // redirect to homepage if user is authenticated
  if (authenticated) {
    console.log("Login successful");

    await router.push('/home');
  }
};
</script>
