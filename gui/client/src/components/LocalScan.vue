<template>
  <div v-if="loading || data.length > 0">
    <h3>{{ name }}</h3>

    <i v-if="loading">Loading...</i>

    <b-table outlined :stacked="data.length == 1" :items="data" v-show="data.length > 0"></b-table>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import axios, { AxiosResponse } from "axios";
import { mapMutations } from "vuex";
import config from "@/config";

interface localScanResponse {
  number: string;
  dork: string;
  URL: string;
}

@Component
export default class GoogleSearch extends Vue {
  id = "local";
  name = "Local scan";
  data: localScanResponse[] = [];
  loading = false;
  computed = {
    ...mapMutations(["pushError"])
  };

  public async created(): Promise<void> {
    this.loading = true;

    try {
      const res = await axios.get(
        `${config.apiUrl}/numbers/13152841580/scan/${this.id}`
      );

      this.data.push(res.data.result);
    } catch (e) {
      this.$store.commit("pushError", { message: e });
    }

    this.loading = false;
  }
}
</script>
