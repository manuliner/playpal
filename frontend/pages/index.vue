<template>
  <VContainer fluid class="fill-height">
    <VCol>
      <VRow no-gutters align="center" justify="center" class="fill-height">
        <VCol class="hidden-md-and-down fill-height" md="6" lg="7">
          <VImg
            src="https://wallpaper.dog/large/5557744.jpg"
            cover
            class="h-100 rounded-xl d-flex align-center justify-center"
          >
            <div class="text-center w-50 text-black mx-auto">
              <h2 class="mb-4">Wilkommen bei playPal</h2>
              <p>Deine Spielplatzverabredungs organisierer</p>
            </div>
          </VImg>
        </VCol>
        <div class="text-center w-50 text-white mx-auto">
          <h2 class="mb-4">Wilkommen bei playPal</h2>
          <p>Deine Spielplatzverabredungs organisierer</p>
        </div>
      </VRow>

      <VRow no-gutters align="center" justify="center" class="fill-height pt-10">
        <VCol cols="12" md="6" lg="5" sm="6">
          <VRow no-gutters align="center" justify="center">
            <VCol>
              <VRow>
                <div class="text-center text-black mx-auto">
                  <h2 class="mb-4">Sign in with email</h2>
                </div>
              </VRow>
              <VRow>
                <div class="text-center text-black mx-auto">
                  <p>
                    Enter the email adress associated with your account and we'll send you a magic
                    link. If you don't have an account, we will create one for you.
                  </p>
                </div>
              </VRow>
              <VRow class="pt-2">
                <VTextField
                  :rules="[ruleRequired, ruleEmail]"
                  v-model="email"
                  prepend-inner-icon="fluent:mail-24-regular"
                  id="email"
                  name="email"
                  type="email"
                />
              </VRow>
              <VRow>
                <VBtn block min-height="44" class="gradient primary" @click="sendMagicLink">
                  Sign In
                </VBtn></VRow
              >
            </VCol>
          </VRow>
          <VRow>{{ authenticated }}</VRow>
        </VCol>
      </VRow>
    </VCol>
  </VContainer>
</template>

<script setup>
const email = ref("");

const router = useRouter();
import { storeToRefs } from "pinia";

import { useAuthStore } from "~/store/auth";
import { useRouter } from "vue-router";
import { VCol, VTextField } from "vuetify/components";

const { authenticated } = storeToRefs(useAuthStore());
const { signInOrUpWithEmail } = useAuthStore();

const { ruleEmail, ruleRequired } = useFormRules();
definePageMeta({
  layout: "default2",
});

const sendMagicLink = async () => {
  console.log("sendMagicLink");
  await signInOrUpWithEmail(email.value);
  if (authenticated.value) {
    console.log("Login successful");

    await router.push("/home");
  }
};
</script>
