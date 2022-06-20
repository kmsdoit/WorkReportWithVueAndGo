<template>
  <!-- NavBar는 로그인 성공한 후 메인페이지에서 보여주자. -->
  <!-- <NavBar/> -->


  <router-view></router-view>


</template>

<script>
// import NavBar from './components/NavBar.vue'

export default {
  name: 'App',
  data() {
    return {

    }
  },
  components: {
    // NavBar,
}
}
</script>

<style>
#app {
  color: #333;  
  font-size: 16px;  
  font-weight: 400; 
  line-height: 1.4; 
  font-family: 'Nanum Gothic', sans-serif;
  /* margin: 0;
  padding: 0;
  box-sizing: border-box; */
  text-align: center;
}
</style>
