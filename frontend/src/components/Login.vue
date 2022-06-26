<template>
  <div class="login-bg">
    <!-- <Header></Header> -->
    <div class="login-content">
      <router-link to="/">
        <img class="logo-img" src="../assets/IDB_logo.png" alt="IDB_LOGO" />
      </router-link>
      <div class="login-box">
        <h2>로그인</h2>
        <form @submit.prevent="doLogin">
          <div class="user-box">
            <label>아이디</label>
            <input
              id="userID"
              type="text"
              name=""
              required=""
              v-model="userId"
            />
          </div>
          <div class="user-box">
            <label>비밀번호</label>
            <input
              id="userPW"
              type="password"
              name=""
              required=""
              v-model="userPw"
            />
          </div>
          <button type="submit" class="login-btn">로그인</button>
        </form>
        <div>
          <!-- <button class="findId-btn">아이디 찾기</button> -->
          <button class="findPw-btn">비밀번호 찾기</button>
          <button class="signUp-btn" @click="goRegister()">회원 가입</button>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
import axios from "axios";
// import Header from './Header.vue';

export default {
  name: "Login",
  data: () => {
    return {
      userId: "",
      userPw: "",
    };
  },
  methods: {
    doLogin() {
      // e.preventDefault();
      let saveData = {};
      saveData.userId = this.userId;
      saveData.userPw = this.userPw;
      // fetch는 비교적 최신 브라우저에서만 동작하므로, 브라우저의 호환성을 위해 axios 라이브러리를 주로 쓴다고 함.
      axios
        .post("url~~~")
        .then((result) => {
          console.log(result);

          // 로그인 성공시
          // this.$router.push("/main");


        })
        .catch((err) => {
          console.log(err);
        });
    },
    goRegister() {
      this.$router.push("/register");
    },
  },
  // components: { Header }
};
</script>

<style>
.login-bg {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  /* background-color: #03a9f4; */
  background-color: #0298db;
  font-size: 16px;
  font-weight: 400;
  line-height: 1.4;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
    Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  text-align: center;
  position: relative;
}

/* 로고 + 로그인 박스 */
.login-content {
  display: block;
  margin: 140px auto;
  width: 100%;
  max-width: 1000px;
  /* height: 400px; */
}

/* logo-img */
.logo-img {
  display: block;
  margin: 0 auto;
  margin-top: 50px;
  margin-bottom: 30px;
  image-rendering: -webkit-optimize-contrast;
  transform: translateZ(0); 
  backface-visibility: hidden;
}

/* login box */
.login-box {
  width: 100%;
  max-width: 600px;
  display: block;
  margin: 0 auto;
  height: 400px;
  background: #fff;
  color: #333;
  border: none;
  border-radius: 10px;
}
.login-box h2 {
  font-size: 26px;
  font-weight: 800;
  padding: 20px;
}

/* input + label in login box */
.user-box {
  width: 100%;
  display: block;
  margin: 10px auto;
  text-align: left;
}
.user-box:last-child {
  margin-top: 30px;
}
.user-box input {
  border: 2px solid #eee;
  border-radius: 5px;
  background-color: #fafafa;
  padding: 10px;
  width: 80%;
  display: block;
  margin: 10px auto;
}
.user-box label {
  width: 80%;
  display: block;
  margin: 20px auto 10px;
}


/* login button */
.login-btn {
  display: inline-block;
  width: 100%;
  max-width: 100px;
  padding: 10px 15px;
  font-size: 20px;
  text-decoration: none;
  letter-spacing: 3px;
  margin-top: 30px;
	color: #fff;
  cursor: pointer;
  transition: 0.5s;
  border: none;
	border-radius: 4px;
  background: #0298db;
}
.login-btn:hover {
  background-color: #0276aa;
}

/* pw,  */
.findPw-btn,
.signUp-btn {
  margin-top: 20px;
  padding: 0 13px 0;
  border: none;
  background: transparent;
  color: #333;
  font-size: 15px;
  transition: 0.5s;
  cursor: pointer;
  position: relative;
}

.findPw-btn::after {
  content: "";
  width: 1px;
  height: 12px;
  background-color: #333;
  position: absolute;
  top: 0;
  bottom: 0;
  right: 0;
  margin: auto;
}

.findPw-btn:hover,
.signUp-btn:hover {
  /* border-bottom: 1px solid #333; */
  color: #0298db;
}
</style>
