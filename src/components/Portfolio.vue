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

      <!-- Custom Tabs -->
      <div class="tabs-container">
        <div class="tabs-header">
          <button 
            class="tab-button" 
            :class="{ 'active': activeTab === 'development' }"
            @click="activeTab = 'development'"
          >
            development
          </button>
          <button 
            class="tab-button" 
            :class="{ 'active': activeTab === 'personal' }"
            @click="activeTab = 'personal'"
          >
            personal projects/42 projects
          </button>
        </div>

        <!-- Development Tab -->
        <div v-if="activeTab === 'development'" class="tab-content">
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
        </div>

        <!-- Personal Projects Tab -->
        <div v-if="activeTab === 'personal'" class="tab-content">
          <div class="row">
            <div
              v-for="(design, idx) in desgin_info"
              :key="idx"
              :class="{ 'mt-4': idx === 0 ? true : true }"
              class="col-xl-6 col-bg-6 col-md-12 col-sm-12"
              style="position: relative;"
            >
              <div style="position: relative;">
                <!-- Simple Carousel -->
                <div class="simple-carousel">
                  <div class="carousel-container">
                    <img 
                      v-for="(slide, i) in design.pictures" 
                      :key="i"
                      :src="slide.img"
                      :class="{ 'active': i === design.currentSlide }"
                      class="carousel-slide"
                      @click="showDesignModalFn(design)"
                    />
                  </div>
                  <div class="carousel-controls">
                    <button 
                      @click="previousSlide(design)"
                      class="carousel-btn"
                      :disabled="design.currentSlide === 0"
                    >
                      ‹
                    </button>
                    <button 
                      @click="nextSlide(design)"
                      class="carousel-btn"
                      :disabled="design.currentSlide === design.pictures.length - 1"
                    >
                      ›
                    </button>
                  </div>
                </div>
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
        </div>
      </div>
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
import info from "../../info";

export default {
  name: "Portfolio",
  components: {
    Card,
    Modal,
    DesignModal,
  },
  props: {
    nightMode: {
      type: Boolean,
    },
  },
  data() {
    return {
      activeTab: 'development',
      all_info: info.portfolio,
      desgin_info: info.portfolio_design.map(design => ({
        ...design,
        currentSlide: 0
      })),
      portfolio_info: info.portfolio,
      showModal: false,
      showDesignModal: false,
      modal_info: {},
      design_modal_info: {},
    };
  },
  methods: {
    nextSlide(design) {
      if (design.currentSlide < design.pictures.length - 1) {
        design.currentSlide++;
      }
    },
    previousSlide(design) {
      if (design.currentSlide > 0) {
        design.currentSlide--;
      }
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

/* Custom Tabs */
.tabs-container {
  margin-top: 20px;
}

.tabs-header {
  display: flex;
  border-bottom: 2px solid #e0e0e0;
  margin-bottom: 20px;
}

.tab-button {
  background: none;
  border: none;
  padding: 10px 20px;
  cursor: pointer;
  font-size: 16px;
  color: #666;
  transition: all 0.3s ease;
  border-bottom: 2px solid transparent;
}

.tab-button:hover {
  color: #669db3ff;
}

.tab-button.active {
  color: #669db3ff;
  border-bottom-color: #669db3ff;
}

.tab-content {
  min-height: 400px;
}

/* Simple Carousel */
.simple-carousel {
  position: relative;
  width: 100%;
  height: 300px;
  border-radius: 10px;
  overflow: hidden;
}

.carousel-container {
  position: relative;
  width: 100%;
  height: 100%;
}

.carousel-slide {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0;
  transition: opacity 0.3s ease;
  cursor: pointer;
}

.carousel-slide.active {
  opacity: 1;
}

.carousel-controls {
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  transform: translateY(-50%);
  display: flex;
  justify-content: space-between;
  padding: 0 10px;
  pointer-events: none;
}

.carousel-btn {
  background: rgba(0, 0, 0, 0.5);
  color: white;
  border: none;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  font-size: 20px;
  cursor: pointer;
  pointer-events: all;
  transition: background 0.3s ease;
}

.carousel-btn:hover:not(:disabled) {
  background: rgba(0, 0, 0, 0.7);
}

.carousel-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.title2 {
  font-size: 20px;
  font-weight: 400;
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

/* Modal transitions */
.modal-enter-active, .modal-leave-active {
  transition: opacity 0.3s;
}

.modal-enter, .modal-leave-to {
  opacity: 0;
}
</style>
