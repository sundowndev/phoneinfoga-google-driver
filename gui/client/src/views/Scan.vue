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
          type="number"
          required
          placeholder="+33678132393"
        ></b-form-input>
      </b-form-group>

      <div>
        <b-button
          variant="outline-primary"
          v-for="scanner in scanners"
          v-bind:key="scanner.name"
          v-on:click="runScan(scanner)"
        >
          {{ scanner.name }}
        </b-button>
      </div>
    </b-form>

    <hr />

    <div v-for="scanner in scanners" v-bind:key="scanner.name" :id="scanner.id">
      <h3>{{ scanner.name }}</h3>
      <p>scan here</p>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState, mapMutations } from "vuex";

interface Scanner {
  id: string;
  name: string;
}

interface ScanData {
  form: {
    number: string;
  };
  scanners: Scanner[];
}

export default Vue.extend({
  computed: {
    ...mapMutations(["pushError"])
  },
  data(): ScanData {
    return {
      form: {
        number: ""
      },
      scanners: [
        { id: "local", name: "Local scan" },
        { id: "numverify", name: "Numverify" },
        { id: "googlesearch", name: "Google search" }
      ]
    };
  },
  methods: {
    async runScan(scanner: Scanner) {
      this.$store.commit("setNumber", this.form.number);

      const res = await this.$store.dispatch("runScanner", scanner.id);
      console.log(res);
    },
    onSubmit(evt: any) {
      evt.preventDefault();
    }
  }
});
</script>
