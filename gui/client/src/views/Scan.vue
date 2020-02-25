<template>
  <div>
    <b-form @submit="onSubmit">
      <b-form-group
        id="input-group-1"
        label="Phone number :"
        label-for="input-1"
        description="Only accepts E164 and International formats as input."
      >
        <b-form-input
          id="input-number"
          v-model="form.number"
          type="text"
          required
          placeholder="e.g: 33678132393"
        ></b-form-input>
      </b-form-group>

      <div>
        <b-button
          variant="outline-primary"
          v-for="scanner in scanners"
          v-bind:key="scanner.name"
          v-on:click="runScan(scanner)"
        >{{ scanner.name }}</b-button>
        <b-button variant="danger" v-on:click="clear">Clear results</b-button>
      </div>
    </b-form>

    <hr />

    <div
      v-for="scanner in scanners"
      v-bind:key="scanner.name"
      :id="scanner.id"
      v-show="scanner.loading || scanner.data.length > 0"
    >
      <h3>{{ scanner.name }}</h3>
      <i v-show="scanner.loading">Running scan for {{ form.number }}...</i>
      <b-table
        outlined
        :stacked="scanner.data.length == 1"
        :items="scanner.data"
        v-show="scanner.data.length > 0"
      ></b-table>
    </div>

    <GoogleSearch/>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { mapMutations } from "vuex";
import GoogleSearch from "../components/GoogleSearch.vue";

interface Scanner {
  id: string;
  name: string;
  data: unknown[];
  loading: boolean;
}

interface ScanData {
  form: {
    number: string;
  };
  scanners: Scanner[];
}

export default Vue.extend({
  components: { GoogleSearch },
  computed: {
    ...mapMutations(["pushError"])
  },
  data(): ScanData {
    return {
      form: {
        number: ""
      },
      scanners: [
        { id: "local", name: "Local scan", data: [], loading: false },
        { id: "numverify", name: "Numverify", data: [], loading: false },
        { id: "googlesearch", name: "Google search", data: [], loading: false }
      ]
    };
  },
  methods: {
    clear() {
      for (const scanner of this.scanners) {
        scanner.loading = false;
        scanner.data = [];
      }
    },
    async runScan(scanner: Scanner): Promise<void> {
      if (this.form.number.length < 2) {
        this.$store.commit("pushError", { message: "Number is not valid." });
        return;
      }

      scanner.loading = true;
      this.$store.commit("setNumber", this.form.number);

      const res = await this.$store.dispatch("runScanner", scanner.id);
      scanner.data.push(res.result);
      scanner.loading = false;
    },
    onSubmit(evt: any) {
      evt.preventDefault();
    }
  }
});
</script>
