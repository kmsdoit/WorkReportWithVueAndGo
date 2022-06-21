<template>
  <nav>
    <div class="inner">

      <!-- 로고 -->
      <a href="/" class="logo link">
        <img src="../assets/IDB_logo.png" alt="IDB_LOGO">
      </a>

      <!-- 메뉴 부분 -->
      <div class="menu-container">
        <ul class="menu">
          <!-- 나중에 메뉴 추가 -->
          <li>
            <button class="login" id="loginBtn" @click="goLogin()">로그인</button>
          </li>
          <li>
            <button class="register" id="registerBtn" @click="goRegister()">회원가입</button>
          </li>
        </ul>
      </div>

    </div>
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
  background-color: #ffffff;
  border-bottom: 1px solid #c8c8c8;
}
img {
  display: block; 
}
li {
  list-style: none;
}
.link {
  text-decoration: none;
}
.inner {    
  width: 80%;
  height: 60px;
  margin: 0 auto;
  position: relative;
}
/* logo */
.logo {
  position: absolute;
  top: 0;
  bottom: 0;
  margin: auto;
  height: 50px; /* logo의 높이값과 같도록 해야한다. */
  left: 0;  /* logo가 inner에 잘 붙어있다고 명시 */
} 
/* menu */
.menu-container {
  display: flex;
  position: absolute;
  top: 5px;
  right: 0;
  background: #eee;
}
.menu-container ul {
  display: flex;
  font-family: Arial, Helvetica, sans-serif;
  font-size: 14px;
}



/* nav .menu-container ul.menu li {
  position: relative;
} */
/* nav .menu-container ul.menu li::before {
  content: "";
  width: 1px;
  height: 12px;
  background-color: #e5e5e5;
  position: absolute; 
  top: 0;
  bottom: 0;
  margin: auto;
} */

/* .menu-container ul.menu li:first-child::before { */
  /* 첫번째 li태그의 가상요소(즉, sing in앞의 구분선)를 안보이도록 */
  /* display: none; 
} */

/* .link {
  font-size: 20px;
  padding: 11px 16px;
  display: block;
  color: #333;
} */

/* .link:hover {
  color: rgb(69, 226, 231);
} */

</style>