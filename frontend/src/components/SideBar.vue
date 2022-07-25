<template>
  <!-- 사이드 바 -->
  <div
    class="sidebar-container"
    :data-color="sidebarItemColor"
    :data-image="sidebarBackgroundImage"
    :style="sidebarStyle"
  >
  <!-- 사이드바 맨위의 로고 -->
    <div class="logo-wrap">
      <a href="https://www.idb.ai" class="logo">
        <img src="../assets/IDB_logo.png" alt="IDB"/>
      </a>
    </div>

    <div class="sidebar-wrap">
      <ul class="sidebar-contents">
        <li class="content1 contents">
          <div>
            <h1><font-awesome-icon icon="fa-solid fa-bicycle" /></h1>
            <p>Home</p>
          </div>
          </li>
        <li class="content2 contents">출퇴근 관리</li>
        <li class="content3 contents">휴가 관리</li>
      </ul>
    </div>
  </div>
</template>
<script>
export default {
  name: 'SideBar',
  data() {
    return {

    }
  },
  // components: {
  //   NavBar,
  // },
  // props: {

  // },
}
</script>
<style>

.sidebar-container {
  background-image: linear-gradient( rgba(0,0,0,0.7), rgba(0,0,0,0.7) ), url('../assets/sidebar-2.jpg');
  background-size: cover;
  background-position: center center;
  color: white;
  z-index: 5;
  width: 300px;
  height: 100vh;
  position: fixed;
  top: 0;
  left: 0;          
}

.logo-wrap {
  width: 100%;
  height: 10%;
}
.logo {
  width: 100%;
  display: flex;
  justify-content: center;
  position: relative;
}
.logo::after {
  content: "";
  position: absolute;
  bottom: 0;
  right: 15px;
  height: 1px;
  width: calc(100% - 30px);
  background-color: rgba(180, 180, 180, 0.3);
}

.sidebar-wrap {
  width: 100%;
  height: 100vh;
  display: flex;
  justify-content: center;
  
}
.sidebar-contents {
  text-align: left;
  display: block;
  margin-top: 20px;
  list-style: none;
  width: 80%;
}
.contents {
  margin-top: 20px;
  widows: 100%;
  padding: 20px;
  transition: all .5s;
  cursor: pointer;
  font-weight: 200;
}
.contents:hover {
  border: none;
  border-radius: 3px;
  background: rgba(200, 200, 200, 0.2);
}
.contents:focus {
  background: #0298db;
}
.contents:active {
  background: #0298db;
}

.content1, .content2, .content3 {

}

</style>