<template>
  <div
    class="py-4 p-st"
    :class="{
      'bg-light': !nightMode,
      'bg-dark2': nightMode,
      'text-light': nightMode,
    }"
  >
    <div class="container">
      <div
        class="text-center"
        data-aos="fade"
        data-aos-once="true"
        data-aos-duration="1000"
      >
        <span
          class="title text-center"
          :class="{ pgray: !nightMode, 'text-light': nightMode }"
          >portfolio.</span
        >
      </div>
      <hr
        width="50%"
        :class="{ pgray: !nightMode, 'bg-secondary': nightMode }"
      />

      <vue-tabs :activeTextColor="!nightMode ? '#535A5E' : '#dfdfdf'">
        <v-tab title="development">
          <br />
          <div class="row">
            <div
              class="col-xl-4 col-bg-4 col-md-6 col-sm-12"
              v-for="(portfolio, idx) in portfolio_info"
              :key="portfolio.name"
            >
              <Card
                :style="{ 'transition-delay': (idx % 3) / 4.2 + 's' }"
                :portfolio="portfolio"
                @show="showModalFn"
                data-aos="fade-up"
                :nightMode="nightMode"
                data-aos-offset="100"
                data-aos-delay="10"
                data-aos-duration="500"
                data-aos-easing="ease-in-out"
                data-aos-mirror="true"
                data-aos-once="true"
              />
            </div>
          </div>
        </v-tab>

        <v-tab title="personal projects/42 projects">
          <div class="row">
            <div
              v-for="(design, idx) in desgin_info"
              :key="idx"
              :class="{ 'mt-4': idx === 0 ? true : true }"
              class="col-xl-6 col-bg-6 col-md-12 col-sm-12"
              style="position: relative;"
            >
              <div style="position: relative;">
                <vueper-slides
                  :dragging-distance="50"
                  fixed-height="300px"
                  :bullets="false"
                  slide-content-outside="bottom"
                  style="position: aboslute"
                    @click.prevent="showDesignModalFn(design)"

                >
                  <vueper-slide
                    v-for="(slide, i) in design.pictures"
                    :key="i"
                    :image="slide.img"
                  />
                </vueper-slides>
                <div 
                  class="collaboration-tag"
                  :class="{
                    'team-tag': design.collaboration === 'Team',
                    'solo-tag': design.collaboration === 'Solo'
                  }"
                >
                  {{ design.collaboration }}
                </div>
              </div>
              <div
                style="width: 100%; display: flex; justify-content: space-between"
                class="mt-2"
              >
                <div>
                  <div class="title2" style="font-weight: 500;">{{ design.title }}</div>
                  <span
                    class="badge mr-2 mb-2"
                    v-for="tech in design.technologies"
                    :key="tech"
                    :class="{ 'bg-dark4': nightMode }"
                    >{{ tech }}</span
                  >
                  •
                  <span class="date ml-1">{{design.date}}</span>
                </div>

                <button
                  style="height: 31px; margin-top: 5px;"
                  class="btn-sm btn btn-outline-secondary no-outline"
                  @click.prevent="showDesignModalFn(design)"
                >
                  read more
                </button>
              </div>
            </div>
          </div>
          <br />
        </v-tab>
      </vue-tabs>
    </div>
    <transition name="modal">
      <Modal
        :showModal="showModal"
        @close="closeModal"
        v-if="showModal"
        :portfolio="modal_info"
        :nightMode="nightMode"
      />
    </transition>
    <transition name="modal">
      <DesignModal
        :showModal="showDesignModal"
        @close="closeModal"
        v-if="showDesignModal"
        :portfolio="design_modal_info"
        :nightMode="nightMode"
      />
    </transition>
  </div>
</template>

<script>
import Card from "./helpers/Card";
import Modal from "./helpers/Modal";
import DesignModal from "./helpers/DesignModal";
import Carousel from "./helpers/Carousel";
import info from "../../info";

import { VueTabs, VTab } from "vue-nav-tabs";
import "vue-nav-tabs/themes/vue-tabs.css";

import { VueperSlides, VueperSlide } from "vueperslides";
import "vueperslides/dist/vueperslides.css";

export default {
  name: "Portfolio",
  components: {
    Card,
    Modal,
    VueTabs,
    VTab,
    VueperSlides,
    VueperSlide,
    DesignModal,
  },
  props: {
    nightMode: {
      type: Boolean,
    },
  },
  data() {
    return {
      all_info: info.portfolio,
      desgin_info: info.portfolio_design,
      portfolio_info: info.portfolio,
      showModal: false,
      showDesignModal: false,
      modal_info: {},
      design_modal_info: {},
      data: [
        '<div class="example-slide">Slide 1</div>',
        '<div class="example-slide">Slide 2</div>',
        '<div class="example-slide">Slide 3</div>',
      ],
    };
  },
  created() {
    // Remove the pagination logic - show all projects
  },
  watch: {
    // Remove the number watcher since we're not using pagination
  },
  methods: {
    next() {
      this.$refs.flickity.next();
    },

    previous() {
      this.$refs.flickity.previous();
    },
    closeModal() {
      this.showModal = false;
      this.showDesignModal = false;
      document.getElementsByTagName("body")[0].classList.remove("modal-open");
    },
    showModalFn(portfolio) {
      this.modal_info = portfolio;
      this.showModal = true;
    },
    showDesignModalFn(design_portfolio) {
      this.design_modal_info = design_portfolio;
      this.showDesignModal = true;
    },
  },
};
</script>

<style scoped>
.title {
  font-size: 30px;
  font-weight: 500;
}
.title1 {
  font-size: 24px;
  font-weight: 400;
}

.title2 {
  font-size: 20px;
  font-weight: 400;
}

.title3 {
  font-size: 16px;
  font-weight: 400;
}

.modal-enter {
  opacity: 0;
}

.modal-leave-active {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave-active .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}

.btn {
  border-color: rgb(212, 149, 97);
  color: rgb(212, 149, 97);
}

.btn:hover {
  background-color: rgb(212, 149, 97);
  border-color: rgb(212, 149, 97);
  color: white;
}

.btn:focus {
  background-color: rgb(212, 149, 97);
  border-color: rgb(212, 149, 97);
  color: white;
}

/deep/ .vue-tabs .nav-tabs {
  border: none;
  font-size: 20px;
  font-weight: 500;
  display: flex;

  justify-content: center;
}

/deep/ .vue-tabs .tabs__link {
  color: #a0a0a0;
}

/deep/ .vue-tabs .nav-tabs > li.active > a {
  background: transparent;
  border: none;
  transition: all 0.5s;
  padding-right: 0;
  padding-left: 0;
  margin-right: 15px;
  margin-left: 15px;
}

/deep/ .vue-tabs .nav-tabs > li > a:hover {
  background: transparent;
  color: #cbcbcb;
  transition: all 0.5s;
}

/deep/ .vue-tabs .nav-tabs > li > a {
  background: transparent;
  border: none;
  transition: all 0.5s;
}

/deep/ .vue-tabs .nav-tabs > li > a:after {
  content: "";
  width: 20%;
  position: absolute;
  bottom: 3px;
  border-width: 0 0 2px;
  border-style: solid;
  transition: all 0.5s;
}

/deep/ .vue-tabs .nav-tabs > li.active > a:after {
  width: 100%;
  transition: all 0.5s;
}

.design-img {
  width: 100%;
  border-radius: 15px;
  transition: all 0.5s;
}

.dimg {
  position: relative;
  border-radius: 15px;
}
.middle {
  transition: all 0.5s;
  opacity: 0;
  position: absolute;
  bottom: 0px;
  left: 70px;
  transform: translate(-50%, -50%);
  -ms-transform: translate(-50%, -50%);
  text-align: center;
  padding: 20px;
}

.dimg:hover .design-img {
  position: relative;
  border-radius: 15px;
  opacity: 0.1;
  cursor: pointer;
}

.dimg:hover .middle {
  opacity: 1;
}

/deep/.vueperslide {
  border-radius: 10px !important;
}
/deep/.vueperslides__parallax-wrapper {
  border-radius: 10px !important;
}

.btn {
  border-color: #669db3ff;
  color: #669db3ff;
}

.btn:hover {
  background-color: #669db3ff;
  border-color: #669db3ff;
  color: white;
}

.btn:focus {
  background-color: #669db3ff;
  border-color: #669db3ff;
  color: white;
}
/deep/ .vueperslides__arrow {
  outline: none !important;
  border: none;
  color: grey;
}

.badge {
  background-color: rgb(211, 227, 233);
  transition: all 0.5s;
  font-weight: 500;
  font-size: 13px;
  padding: 2px 8px;
}

.bg-dark4 {
  background-color: #494e55 !important;
}

.date {
  font-size: 14px;
  font-weight: 400;
  opacity: 0.75
}

.collaboration-tag {
  position: absolute;
  top: 8px;
  right: 8px;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: white;
  z-index: 10;
}

.team-tag {
  background-color: #28a745;
}

.solo-tag {
  background-color: #007bff;
}
</style>
