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
          v-model="number"
          type="text"
          required
          placeholder="e.g: 33678132393"
          :disabled="loading"
        ></b-form-input>
      </b-form-group>

      <b-button size="sm" variant="dark" v-on:click="runScans" :disabled="loading">
        <b-icon-play-fill></b-icon-play-fill>Run scan
      </b-button>

      <b-button variant="light" size="sm" v-on:click="clearData">Clear results</b-button>
    </b-form>

    <hr />

    <LocalScan />
    <NumverifyScan />
    <GoogleSearch />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { mapMutations } from "vuex";
import LocalScan from "../components/LocalScan.vue";
import NumverifyScan from "../components/NumverifyScan.vue";
import GoogleSearch from "../components/GoogleSearch.vue";

interface Scanner {
  id: string;
  name: string;
  data: object[];
  loading: boolean;
}

interface Data {
  loading: boolean;
  number: string;
}

export default Vue.extend({
  components: { LocalScan, GoogleSearch, NumverifyScan },
  computed: {
    ...mapMutations(["pushError"])
  },
  data(): Data {
    return {
      loading: false,
      number: ""
    };
  },
  methods: {
    clearData() {
      console.log("clear data");
    },
    // async runScan(scanner: Scanner): Promise<void> {
    //   if (this.number.length < 2) {
    //     this.$store.commit("pushError", { message: "Number is not valid." });
    //     return;
    //   }

    //   scanner.loading = true;
    //   this.$store.commit("setNumber", this.number);

    //   const res = await this.$store.dispatch("runScanner", scanner.id);
    //   scanner.data.push(res.result);
    //   scanner.loading = false;
    // },
    async runScans(scanner: Scanner): Promise<void> {
      if (this.number.length < 2) {
        this.$store.commit("pushError", { message: "Number is not valid." });
        return;
      }

      this.loading = true;

      this.$store.commit("setNumber", this.number);

      this.loading = false;
    },
    onSubmit(evt: any) {
      evt.preventDefault();
    }
  }
});
</script>
