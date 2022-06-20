<template>
  <div class='login-container'>
    <div class="login-box">
      <h2>로그인</h2>
        <form @submit.prevent="doLogin">
          <div class="user-box">
            <input id="userID" type="text" name="" required="" v-model="userId"/>
            <label>이메일</label>
          </div>
          <div class="user-box">
            <input id="userPW" type="password" name="" required="" v-model="userPw"/>
            <label>비밀번호</label>
          </div>
          <button type="submit" class='login-btn'>로그인</button>
        </form>
      <div class="register">
        <button class='findId-btn'>아이디 찾기</button>
        <button class='findPw-btn'>비밀번호 찾기</button>
        <button class='signUp-btn' @click="goRegister()">회원 가입</button>
      </div>
    </div>
  </div>
</template>

<script>

import axios from 'axios';

export default {
  name: 'Login',
  data: () => {
    return {
      userId: '',
      userPw: '',

    }
  },
  methods: {
    doLogin() {
      // e.preventDefault();

      let saveData = {};
      saveData.userId = this.userId;
      saveData.userPw = this.userPw;

    // fetch는 비교적 최신 브라우저에서만 동작하므로, 브라우저의 호환성을 위해 axios 라이브러리를 주로 쓴다고 함. 
      axios.post('url~~~')
        .then( result => {
          console.log(result);



        }).catch((err) => {
          console.log(err)
      })

      
    },
    goRegister() {
      this.$router.push('/register');
    }

  },



}
</script>

<style>
.login-container {
  width: 100vw;
  height: 100vh;
  overflow-x: hidden;
  overflow-y: auto;
  margin: 0;
  padding: 0;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
    Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
  box-sizing: border-box;
  /* background: linear-gradient(#141e30, #243b55); */
  color: white;
  font-size: 16px;
  font-weight: 400; 
  line-height: 1.4; 
}

.login-box {
  display: block;
  margin: 150px auto;
  width: 400px;
  /* height: 350px; */
  padding: 40px;
  background: rgba(0, 0, 0, 0.5);
  border-radius: 10px;
  box-shadow: 0 15px 25px rgba(0, 0, 0, 0.6);

  /* transform: translate(-50%, -50%); */
}

.login-box h2 {
  margin: 0 0 30px;
  padding: 0;
  text-align: center;
}

.user-box {
  position: relative;
}


.user-box input {
  width: 100%;
  padding: 10px 0;
  font-size: 16px;
  color: #fff;
  margin-bottom: 30px;
  border: none;
  border-bottom: 1px solid #fff;
  outline: none;
  background: transparent;
}
.user-box label {
  position: absolute;
  top: 0;
  left: 0;
  padding: 10px 0;
  font-size: 16px;
  color: #fff;
  pointer-events: none;
  transition: 0.5s;
}

.user-box input:hover ~ label, .user-box input:valid ~ label {
  top: -20px;
  left: 0;
  color: #e5e5e5;
  font-size: 14px;
}

.user-box input:focus ~ label, .user-box input:valid ~ label {
  top: -20px;
  left: 0;
  color: #c8c8c8;
  font-size: 14px;
}

/* 클릭했을때 클래스 붙이기  */
/* .labelUp {
  top: -20px;
  left: 0;
  color: #03e9f4;
  font-size: 12px;
} */


.login-btn {
  position: relative;
  display: inline-block;
  width: 100%;
  padding: 10px 15px;
  color: white;
  font-size: 20px;
  text-decoration: none;
  /* text-transform: uppercase;  대문자 */
  /* overflow: hidden; */
  transition: 0.5s;
  letter-spacing: 3px;
  margin-top: 30px;
  border: 1px solid #fff;
  background: transparent;
  cursor: pointer;
}
.login-btn:hover {
  color: #333;
  border-radius: 5px;
  background-color: #c8c8c8;
}

.register {
  display: flex;
  justify-content: center;
  align-items: center;
  /* width: 400px; */
  text-align: center;
}

.findId-btn,
.findPw-btn,
.signUp-btn {
  margin-top: 50px;
  padding: 0 13px 0;
  border: none;
  background: transparent;
  color: white;
  font-size: 15px;
  transition: 0.5s;
  cursor: pointer;
  position: relative;
}
.findId-btn::after, .findPw-btn::after {
  content: "";
  width: 1px;
  height: 12px;
  background-color: #e5e5e5;
  position: absolute; 
  top: 0;
  bottom: 0;
  right: 0;
  margin: auto;
}

.findId-btn:hover,
.findPw-btn:hover,
.signUp-btn:hover {
  /* cursor: pointer; */
  border-bottom: 1px solid #fff;
  color: #333;
}


</style>