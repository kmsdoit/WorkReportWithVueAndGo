<template>
  <div class="window-container">
    <!-- 사이드바 컨테이너-->
    <div class="sidebar-container">
      <!-- 사이드바 맨위의 로고 -->
      <div class="logo-wrap">
        <a href="http://idb.ai/" class="logo">
          <img src="../assets/IDB_logo.png" alt="IDB" />
        </a>
      </div>
      <!-- 사이드바 내용 -->
      <div class="sidebar-wrap">
        <ul class="sidebar-contents">
          <li class="content1 contents" id="home" @click="showHome()">
            <div>
              <h1 class="sidebar-icon"><font-awesome-icon icon="fa-solid fa-house" size="sm" /></h1>
              <p class="sidebar-home">HOME</p>
            </div>
          </li>
          <li class="content2 contents" id="work" @click="showWork()">
            <div>
              <h1 class="sidebar-icon"><font-awesome-icon icon="fa-solid fa-business-time" size="sm"/></h1>
              <p class="sidebar-work">출퇴근 관리</p>
            </div>
          </li>
          <li class="content3 contents" id="vacation" @click="showVacation()">
            <!-- <router-link to="/main/vacation"> -->
            <div>
              <h1 class="sidebar-icon"><font-awesome-icon icon="fa-solid fa-calendar-check" /></h1>
              <p class="sidebar-vacation">휴가 관리</p>
            </div>
            <!-- </router-link> -->
          </li>
        </ul>
      </div>
    </div>

    <div class="dashboard-container">
      <NavBar
        :home="home"
        :work="work"
        :vacation="vacation"
        :step="step"
      ></NavBar>
      <div v-if="step == 0">
        <Home></Home>
      </div>
      <div v-if="step == 1">
        <Work></Work>
      </div>
      <div v-if="step == 2">
        <Vacation></Vacation>
      </div>
    </div>

    <!-- home, work, vacation 컴포넌트가 와야함 -->
  </div>
</template>

<script>
// import NavBar from './components/NavBar.vue'
import Home from "../components/Home.vue";
import Work from "../components/Work.vue";
import Vacation from "../components/Vacation.vue";
import NavBar from "../components/NavBar.vue";

export default {
  name: "Main",
  data() {
    return {
      home: "HOME",
      work: "출퇴근 관리",
      vacation: "휴가 관리",
      step: 0,
    };
  },
  components: { Home, Work, Vacation, NavBar },
  methods: {
    showHome() {
      document.getElementById('work').classList.remove('click-bg')
      document.getElementById('vacation').classList.remove('click-bg')
      document.getElementById('home').classList.add('click-bg')
      this.step = 0;
    },
    showWork() {
      document.getElementById('home').classList.remove('click-bg')
      document.getElementById('vacation').classList.remove('click-bg')
      document.getElementById('work').classList.add('click-bg')
      this.step = 1;
    },
    showVacation() {
      document.getElementById('home').classList.remove('click-bg')
      document.getElementById('work').classList.remove('click-bg')
      document.getElementById('vacation').classList.add('click-bg')
      this.step = 2;
    },
    
  },
  mounted() {
    document.getElementById('work').classList.remove('click-bg')
    document.getElementById('vacation').classList.remove('click-bg')
    document.getElementById('home').classList.add('click-bg')
  },
};
</script>

<style>
.window-container {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  background: #eaf2f7;
  position: relative;
}

/*^ 사이드바 */
.sidebar-container {
  background-image: linear-gradient(rgba(0, 0, 0, 0.7), rgba(0, 0, 0, 0.5)),
    url("../assets/sidebar-2.jpg");
  background-size: cover;
  background-position: center center;
  color: white;
  z-index: 5;
  height: 100vh;
  position: absolute;
  top: 0;
  left: 0;
  box-shadow: 2px 0px 10px 1px gray;
  width: 17%;
}
/* 로고 */
.logo-wrap {
  width: 100%;
  height: 100px;
  padding: 20px;
  position: relative;
}
.logo {
  width: 100%;
  display: flex;
  justify-content: center;
  align-content: center;
  position: relative;
}
/* 로고아래 실선 */
.logo-wrap::after {
  content: "";
  position: absolute;
  bottom: 0;
  right: 15px;
  height: 1px;
  width: calc(100% - 30px);
  background-color: rgba(180, 180, 180, 0.3);
}

/* 로고 아래 내용 */
.sidebar-wrap {
  width: 100%;
  height: 100vh;
}
.sidebar-contents {
  text-align: left;
  display: block;
  margin: 20px auto 0;
  list-style: none;
  width: 80%;
}
.contents {
  margin-top: 20px;
  width: 100%;
  padding: 20px;
  transition: all 0.5s;
  cursor: pointer;
  font-weight: 200;
}
.click-bg {
  background: #0298db;
  border-radius: 5px;
}
.contents:hover {
  border: none;
  border-radius: 3px;
  background: rgba(200, 200, 200, 0.2);
}
.contents:focus {
  background: #0298db;
}
/* .contents:active {
  background: #0298db;
} */

.contents div {
  display: flex;
  align-items: flex-start;
  width: 100%;
}
/* 아이콘 */
.sidebar-icon {
  width: 20%;
  margin-left: 20px;
  font-size: 18px;
}
/* 텍스트 */
.sidebar-home,
.sidebar-work,
.sidebar-vacation {
  width: 100%;
  text-align: left;
  margin-left: 15px;
  font-size: 16px;
  margin-top: 2px;    /* font-size 맞추기 */
}

/* 대쉬보드 */
.dashboard-container {
  position: absolute;
  right: 0;
  width: 83%;
  /* height: 100vh;  */
}
</style>
