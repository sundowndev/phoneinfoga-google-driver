import axios, { AxiosResponse } from "axios";
import Vue from "vue";
import Vuex, { Store } from "vuex";
import config from "@/config";

Vue.use(Vuex);

interface Error {
  message: string;
}

interface StoreInterface {
  number: string;
  errors: Error[];
}

const store: Store<StoreInterface> = new Vuex.Store({
  state: {
    number: "",
    errors: [] as Error[]
  },
  mutations: {
    pushError(state, err): void {
      state.errors.push(err);
    },
    setNumber(state, number): void {
      state.number = number;
    },
    resetState(state) {
      state.number = "";
      state.errors = [];

      return state;
    }
  },
  getters: {},
  actions: {
    async runScanner(
      context,
      scanner: string
    ): Promise<AxiosResponse<unknown> | void> {
      try {
        const res = await axios.get(
          `${config.apiUrl}/numbers/${context.state.number}/scan/${scanner}`
        );

        return res.data;
      } catch (e) {
        context.commit("pushError", { code: "ok", message: e });
      }
    },
    resetState(context): void {
      context.commit("resetState");
    }
  },
  modules: {}
});

export default store;
