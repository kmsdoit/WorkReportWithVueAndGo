<template>
<!-- background -->
  <div class="background">
    <Header></Header>
    <!-- 컨텐츠 -->
    <div class="explain-box">
      <h1>IDB 출퇴근 관리</h1>
      <p>IDB 출퇴근 관리 페이지입니다.</p>
    </div>
    <!-- 버튼 -->
    <div class="btn-box">
      <button type="button" class="login-move-btn" @click="goLogin()" >로그인</button>
      <button type="button" class="register-move-btn" @click="goRegister()" >회원가입</button>
    </div>
    <router-view></router-view>
  <!-- <footer class="footer">
    <div class="footer-inner">
      1. #333 2. 유튜브 트위터 페이스북 등
    </div>
  </footer> -->
  </div>
  
</template>

<script>
import Header from './Header.vue';
export default {
  name: "StartView",
  data() {
      return {};
  },
  methods: {
      goLogin() {
          this.$router.push("/login");
      },
      goRegister() {
          this.$router.push("/register");
      },
  },
  components: {
    Header
  }
};
</script>

<style>
.background {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  font-size: 16px;
  font-weight: 400;
  line-height: 1.4;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  text-align: center;
  position: relative;
}

/* explain-box */
.explain-box {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 60%;
  position: absolute;
  top: 80px;
  color: #fff;
}
.explain-box h1 {
  font-size: 50px;
  font-weight: 700;
  line-height: 2;
  /* font-family: 나중에 글꼴 추가; */
}
.explain-box p {
  font-size: 17px;
} 

/* btn-box */
.btn-box {
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  bottom: 40vh;
  left: 0;
  right: 0;
  margin-top: 50px;
}
.btn-box button {
  width: 150px;
  height: 60px;
  border: none;
  border-radius: 10px;
  padding: 10px;
  margin: 15px;
  color: #fff;
  background-color: #0298db;
  font-size: 20px;
  letter-spacing: 2px;
  cursor: pointer;
  transition: .3s;
}
.btn-box button:hover {
  background-color: #0276aa;
  /* color: #0298db; */
}

/* footer */
/* .footer {
  width: 100vw;
  height: 20vh;
  overflow: hidden;
  background-color: #eaf2f7;
  position: absolute;
  bottom: 0;
} */
  

</style>
