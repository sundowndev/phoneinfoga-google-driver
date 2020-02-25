<template>
  <div v-if="data.length > 0">
    <h3>{{ name }}</h3>

    <b-list-group>
      <b-list-group-item :href="value" target="blank" v-for="(value, i) in data" v-bind:key="i">
        {{ value }}
      </b-list-group-item>
    </b-list-group>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import axios, { AxiosResponse } from "axios";
import config from "@/config";

@Component
export default class GoogleSearch extends Vue {
  id = "googlesearch";
  name = "Google search";
  loading = false;
  data: string[] = [];

  async runScan(): Promise<AxiosResponse<unknown> | void> {
    try {
      const res = await axios.get(
        `${config.apiUrl}/numbers/${"112"}/scan/${this.id}`
      );

      return res.data;
    } catch (e) {
      // context.commit('pushError', { message: e });
    }
  }
}
</script>
