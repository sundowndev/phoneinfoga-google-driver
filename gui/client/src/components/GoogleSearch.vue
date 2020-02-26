<template>
  <div v-if="loading || data.length > 0">
    <h3>
      {{ name }}
      <small>
        <b-button variant="outline-primary" size="sm" v-on:click="openLinks"
          >Open all links</b-button
        >
      </small>
    </h3>

    <i v-if="loading">Loading...</i>

    <b-button
      size="sm"
      variant="dark"
      v-b-toggle.googlesearch-collapse
      v-show="data.length > 0"
      >Toggle results</b-button
    >
    <b-collapse id="googlesearch-collapse" class="mt-2">
      <b-list-group>
        <b-list-group-item
          :href="value.URL"
          target="blank"
          v-for="(value, i) in data"
          v-bind:key="i"
          >{{ value.dork }}</b-list-group-item
        >
      </b-list-group>
    </b-collapse>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import axios, { AxiosResponse } from "axios";
import { mapMutations } from "vuex";
import config from "@/config";

interface GoogleSearchScanResponse {
  number: string;
  dork: string;
  URL: string;
}

@Component
export default class GoogleSearch extends Vue {
  id = "googlesearch";
  name = "Google search";
  data: GoogleSearchScanResponse[] = [];
  loading = false;
  computed = {
    ...mapMutations(["pushError"])
  };

  public async created(): Promise<void> {
    this.loading = true;

    try {
      const res: AxiosResponse = await axios.get(
        `${config.apiUrl}/numbers/13152841580/scan/${this.id}`
      );

      this.data = res.data.result;
    } catch (e) {
      this.$store.commit("pushError", { message: e });
    }

    this.loading = false;
  }

  openLinks(): void {
    for (const result of this.data) {
      window.open(result.URL, "_blank");
    }
  }
}
</script>
