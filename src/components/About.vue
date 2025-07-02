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
          >about me.</span
        >
      </div>
      <hr
        width="50%"
        :class="{ pgray: !nightMode, 'bg-secondary': nightMode }"
      />
      <div class="row">
        <div class="col-xl-6 col-bg-6 col-md-6 col-sm-12">
          <Timeline :data="education" :nightMode="nightMode" />
        </div>
        <div class="col-xl-6 col-bg-6 col-md-6 col-sm-12">
          <Timeline :data="experience" :nightMode="nightMode" />
        </div>
      </div>
      
      <!-- Language Skills Section -->
      <div class="row mt-5">
        <div class="col-12">
          <div
            class="text-center mb-4"
            data-aos="fade"
            data-aos-once="true"
            data-aos-duration="1000"
          >
            <span
              class="title text-center"
              :class="{ pgray: !nightMode, 'text-light': nightMode }"
              >language skills.</span
            >
          </div>
          <hr
            width="50%"
            :class="{ pgray: !nightMode, 'bg-secondary': nightMode }"
          />
          <div class="row justify-content-center">
            <div
              v-for="language in languages"
              :key="language.name"
              class="col-xl-4 col-bg-4 col-md-6 col-sm-12 mb-4"
            >
              <div
                class="language-card"
                :class="{
                  'language-card-dark': nightMode,
                  'language-card-light': !nightMode,
                }"
              >
                <h5 class="language-name">{{ language.name }}</h5>
                <div class="language-flag">{{ language.flag }}</div>
                <div class="rating-container">
                  <div
                    v-for="star in 5"
                    :key="star"
                    class="rating-star"
                    :class="{
                      'filled': star <= language.rating,
                      'empty': star > language.rating,
                      'star-dark': nightMode
                    }"
                  ></div>
                </div>
                <span class="language-tag">{{ language.tag }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Timeline from "./helpers/Timeline";
import info from "../../info";

export default {
  name: "About",
  components: {
    Timeline,
  },
  props: {
    nightMode: {
      type: Boolean,
    },
  },
  data() {
    return {
      education: {
        title: "education",
        data: info.education,
      },
      experience: {
        title: "experiences",
        data: info.experience,
      },
      languages: info.languages,
    };
  },
};
</script>

<style scoped>
.title {
  font-size: 30px;
  font-weight: 500;
}

.language-card {
  padding: 20px;
  border-radius: 10px;
  text-align: center;
  transition: all 0.3s ease;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.language-card-light {
  background-color: white;
  border: 1px solid #e0e0e0;
}

.language-card-dark {
  background-color: #30363a;
  border: 1px solid #404040;
}

.language-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.15);
}

.language-name {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 8px;
  color: inherit;
}

.language-flag {
  font-size: 32px;
  margin-bottom: 15px;
  line-height: 1;
}

.rating-container {
  display: flex;
  justify-content: center;
  gap: 8px;
  margin-bottom: 15px;
}

.rating-star {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  transition: all 0.3s ease;
}

.rating-star.filled {
  background-color: #669db3ff;
}

.rating-star.empty {
  background-color: #e0e0e0;
  border: 2px solid #d0d0d0;
}

.rating-star.filled.star-dark {
  background-color: #669db3ff;
}

.rating-star.empty.star-dark {
  background-color: #404040;
  border: 2px solid #505050;
}

.language-tag {
  display: inline-block;
  padding: 4px 12px;
  background-color: #669db3ff;
  color: white;
  border-radius: 15px;
  font-size: 12px;
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}
</style>
