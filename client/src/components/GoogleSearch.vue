<template>
  <div v-if="loading || data.length > 0">
    <h3>
      {{ name }}
      <b-spinner v-if="loading" type="grow"></b-spinner>
      <small>
        <b-button variant="outline-primary" size="sm" v-on:click="openLinks"
          >Open all links</b-button
        >
      </small>
    </h3>

    <b-button
      size="sm"
      variant="dark"
      v-b-toggle.googlesearch-collapse
      v-show="data.length > 0 && !loading"
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
import { Component, Vue, Prop } from "vue-property-decorator";
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

  @Prop() scan: Vue;

  mounted() {
    this.scan.$on("scan", this.run);
    this.scan.$on("clear", this.clear);
  }

  private clear() {
    this.data = [];
  }

  private async run(): Promise<void> {
    this.loading = true;

    try {
      const res: AxiosResponse = await axios.get(
        `${config.apiUrl}/numbers/${this.$store.state.number}/scan/${this.id}`
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
