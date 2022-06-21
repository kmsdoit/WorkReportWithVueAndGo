<template>
  <div class="background">
    <header class="header">
      <img src="./assets/IDB_logo.png" alt="IDB_LOGO"/>
    </header>
    <div class="title">
      <h1>IDB 출퇴근 관리</h1>
    </div>
    <div class="content">
      <p>IDB 출퇴근 관리 페이지 입니다.</p>
    </div>
    <div>
      로그인버튼, 회원가입버튼
    </div>
  </div>
  <footer class="footer">1. #333 2. 유튜브 트위터 페이스북 등</footer>
</template>

<script>
// import NavBar from './components/NavBar.vue'

export default {
  name: "App",
  data() {
    return {};
  },
  components: {
    // NavBar,
  },
};
</script>

<style>
.background {
  width: 100vw;
  height: 80vh;
  overflow: hidden;
  /* background-image: url("./assets/background.jpg"); */
  /* background-size: cover;
  background-repeat: no-repeat;
  background-position: center; */
  background-color: #eaf2f7;

  color: #333;
  font-size: 16px;
  font-weight: 400;
  line-height: 1.4;
  font-family: "Nanum Gothic", sans-serif;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  text-align: center;
}

.header {
  width: 80%;
  height: 60px;
  margin: 0 auto;
  position: relative;
}
img {
  display: block;
  margin: 10px 0;
}
.footer {
  width: 80vw;
  margin: 0 auto;
  height: 20vh;
  overflow: hidden;
  border: none;
  border-top: 1px solid #333;
  background-color: #eaf2f7;
  /* color: #fff; */
}

  

</style>
