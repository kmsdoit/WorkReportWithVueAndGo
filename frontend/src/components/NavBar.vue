<template>
  <nav class="nav-container">
    
      <div class="nav-title">
        <h3 v-if="step == 0">{{home}}</h3>
        <h3 v-if="step == 1">{{work}}</h3>
        <h3 v-if="step == 2">{{vacation}}</h3>
      </div>

      <ul class="nav-menu">
        <li class="menu">
          <button class="login-btn2" id="loginBtn" @click="goLogin()">로그인</button>
        </li>
        <li class="menu">
          <button class="register-btn2" id="registerBtn" @click="goRegister()">회원가입</button>
        </li>
      </ul>
      
  </nav>
</template>

<script>


export default {
  name: 'NavBar',
  data: () => {
    return {

    }
  },
  methods: {
    goLogin() {
      this.$router.push('/login');
    },
    goRegister() {
      this.$router.push('/register');
    }
  },
  props: {
    home: String,
    work: String,
    vacation: String,
    step: Number,
  }
}
</script>

<style>
/* common */
nav {
  color: #333; 
  font-size: 16px;  
  font-weight: 400; 
  line-height: 1.4; 
  font-family: 'Nanum Gothic', sans-serif;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  /* background: #989898; */
  margin: 0 auto;
  width: 100%;
  height: 60px;
  position: relative;
}
li {
  list-style: none;
}

/* 왼쪽 타이틀 */
.nav-title {
  position: absolute;
  left: 0;
  display: flex;
  width: 150px;
  height: 60px;
  justify-content: center;
  align-items: center;
  letter-spacing: 2px;
  font-size: 20px;
  font-weight: lighter;
  /* font-family: ; */
}

/* 오른쪽 메뉴 */
.nav-menu {
  position: absolute;
  right: 0;
  display: flex;
  /* background: #000; */
}
.nav-menu li button {
  padding: 10px;
  margin: 5px 10px 5px;
  background: transparent;
  color: #333;
  border: none;
  transition: .3s;
}
.login-btn2:hover, .register-btn2:hover {
  color: #0298db;
  cursor: pointer;
}

</style>