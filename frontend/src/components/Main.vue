<template>
  <div>
    <SideBar></SideBar>
  </div>

</template>

<script>
// import NavBar from './components/NavBar.vue'
import SideBar from './SideBar.vue';

export default {
    name: "Main",
    data() {
        return {};
    },
    components: { SideBar }
}
</script>

<style>
/* #app {
  color: #333;  
  font-size: 16px;  
  font-weight: 400; 
  line-height: 1.4; 
  font-family: 'Nanum Gothic', sans-serif;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  text-align: center;
} */
</style>
