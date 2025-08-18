<template>
  <div>
    <nav
      class="navbar navbar-expand-lg navbar-light fixed-top p-st"
      :class="{
        'bg-light': !nightMode,
        'navbar-blur': navbarConfig.blur,
        'bg-dark2': nightMode,
      }"
    >
      <div class="container">
        <a
          class="navbar-brand"
          href="/"
          @click.prevent="$emit('scroll', 'home')"
        >
          <Logo :nightMode="nightMode" />
        </a>
        <button
          class="navbar-toggler"
          type="button"
          data-toggle="collapse"
          data-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span style="color: gray; font-size: 23px;"
            ><i class="fas fa-bars"></i
          ></span>
        </button>

        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav ml-auto">
            <li class="nav-item mx-2">
              <a
                class="nav-link"
                href="/about"
                @click.prevent="$emit('scroll', 'about')"
                :class="{ 'text-light': nightMode }"
                >{{ translations.navbar.about }}</a
              >
            </li>
            <li class="nav-item mx-2">
              <a
                class="nav-link"
                href="/skills"
                @click.prevent="$emit('scroll', 'skills')"
                :class="{ 'text-light': nightMode }"
                >{{ translations.navbar.skills }}</a
              >
            </li>
            <li class="nav-item mx-2 ">
              <a
                class="nav-link"
                href="/portfolio"
                @click.prevent="$emit('scroll', 'portfolio')"
                :class="{ 'text-light': nightMode }"
                >{{ translations.navbar.portfolio }}</a
              >
            </li>
            <li class="nav-item mx-2">
              <a
                class="nav-link"
                href="mailto:ryanbwgt@gmail.com"
                :class="{ 'text-light': nightMode }"
                ><i class="fas fa-envelope"></i></a
              >
            </li>
            <li class="nav-item mx-2">
              <div class="language-switcher">
                <button
                  class="btn btn-sm language-btn"
                  :class="{ 
                    'active': currentLanguage === 'en',
                    'btn-outline-light': nightMode,
                    'btn-outline-dark': !nightMode
                  }"
                  @click="setLanguage('en')"
                >
                  EN
                </button>
                <button
                  class="btn btn-sm language-btn"
                  :class="{ 
                    'active': currentLanguage === 'de',
                    'btn-outline-light': nightMode,
                    'btn-outline-dark': !nightMode
                  }"
                  @click="setLanguage('de')"
                >
                  DE
                </button>
              </div>
            </li>
            <li class="nav-item ml-2">
              <a
                class="nav-link"
                href="#"
                @click.prevent="switchMode"
                :class="{ 'text-light': nightMode }"
                ><i
                  :class="{
                    'fas fa-moon': nightMode,
                    'far fa-moon': !nightMode,
                  }"
                  v-tooltip.bottom="nightMode ? 'Light Mode' : 'Night Mode'"
                ></i
              ></a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  </div>
</template>

<script>
import Logo from "./helpers/Logo";
import languageService from "../services/languageService";

export default {
  name: "Navbar",
  props: {
    nightMode: {
      type: Boolean,
    },
  },
  data() {
    return {
      navbarConfig: languageService.getInfo().config.navbar,
      localNightMode: this.nightMode,
      currentLanguage: languageService.getCurrentLanguage(),
      translations: languageService.getTranslations(),
    };
  },
  components: {
    Logo,
  },
  created() {
    // Subscribe to language changes
    languageService.subscribe(this.onLanguageChange);
    // Initialize language service
    languageService.init();
  },
  beforeDestroy() {
    // Unsubscribe when component is destroyed
    languageService.unsubscribe(this.onLanguageChange);
  },
  methods: {
    switchMode() {
      this.localNightMode = !this.localNightMode;
      this.$emit("nightMode", this.localNightMode);
    },
    setLanguage(language) {
      languageService.setLanguage(language);
    },
    onLanguageChange(language, info) {
      this.currentLanguage = language;
      this.navbarConfig = info.config.navbar;
      this.translations = languageService.getTranslations();
      // Emit language change to parent components
      this.$emit("languageChange", language, info);
    },
  },
};
</script>

<style scoped>
.nav-link {
  font-weight: 500;
}

button {
  border: none;
  outline: none;
}

button:hover {
  border: none;
  outline: none;
}

nav {
  border-bottom: 1px solid rgba(160, 159, 159, 0.336);
  position: fixed !important;
}

.navbar-blur {
  background-color: #ffffff7e;
  backdrop-filter: blur(12px);
}

.language-switcher {
  display: flex;
  gap: 2px;
}

.language-btn {
  font-size: 12px;
  padding: 4px 8px;
  font-weight: 600;
  border-radius: 4px;
  transition: all 0.3s ease;
}

.language-btn.active {
  background-color: #669db3ff !important;
  border-color: #669db3ff !important;
  color: white !important;
}

.language-btn:hover {
  background-color: #669db3ff !important;
  border-color: #669db3ff !important;
  color: white !important;
}
</style>
