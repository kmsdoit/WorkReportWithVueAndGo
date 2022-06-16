<template>
  <nav>
    <div class="inner">
      <!-- 로고 -->
      <router-link to="/" class="logo link">
        <img src="../assets/workLogo.png" alt="출퇴근">
      </router-link>

      <div class="menu-container">
        <ul class="menu">
          <li>
            <router-link to="/login" class="login link">로그인</router-link>
          </li>
          <li>
            <router-link to="/signup" class="sign-up link">회원가입</router-link>
          </li>
          <!-- <li>
            <routerLink href="/mypage"></routerLink>
          </li> -->
        </ul>
      </div>

    </div>
  </nav>
</template>

<script>
export default {
  name: 'NavBar',

}
</script>

<style>
/* common */
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
  width: 1300px;
  margin: 0 auto;
  position: relative;
}

/* nav */
nav { 
  /* background-color: #f6f5f0; */
  background-color: rgba(0, 0, 0, 0.5);
  border-bottom: 1px solid #c8c8c8; 
}
nav > .inner {   
  height: 90px;
}
.logo {
  position: absolute;
  top: 0;
  bottom: 0;
  margin: auto;
  height: 75px; /* logo의 높이값과 같도록 해야한다. */
  left: 0;  /* logo가 inner에 잘 붙어있다고 명시 */
}  
nav .menu-container {
  display: flex;
  position: absolute;
  top: 3px;
  right: 0;
}
nav .menu-container ul.menu {
  display: flex;
  font-family: Arial, Helvetica, sans-serif;
}
nav .menu-container ul.menu li {
  position: relative;
}
nav .menu-container ul.menu li::before {
  content: "";
  width: 1px;
  height: 12px;
  background-color: #e5e5e5;
  position: absolute; 
  top: 0;
  bottom: 0;
  margin: auto;
}
nav .menu-container ul.menu li:first-child::before {
  /* 첫번째 li태그의 가상요소(즉, sing in앞의 구분선)를 안보이도록 */
  display: none; 
}
nav .menu-container ul.menu li .link {
  font-size: 20px;
  padding: 11px 16px;
  display: block;
  /* color: #656565; */
  color: #fff;
}
nav .menu-container ul.menu li .link:hover {
  color: #000;
}

</style>