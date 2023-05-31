import { defineStore } from "pinia";
import { useStorage } from "@vueuse/core";
import { setLocale } from "@/i18n";

const useAppStore = defineStore("app", {
  state: () => ({
    language: useStorage("language", "zh"),
  }),
  getters: {},
  actions: {
    setLanguage(lang: string) {
      this.language = lang;
      setLocale(lang);
    },
  },
});

export default useAppStore;
