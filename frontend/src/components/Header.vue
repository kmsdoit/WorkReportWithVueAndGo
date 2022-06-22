<template>
    <!-- 헤더 컴포넌트 -->
    <header class="header">
      <router-link to="/">
        <img src="../assets/IDB_logo.png" alt="IDB_LOGO"/>
      </router-link>
    </header>

</template>

<script>
export default {
  name: "Header",
  data() {
    return {

    };
  },
};
</script>

<style>
/* header */
.header {
  width: 80%;
  height: 60px;
  margin: 0 auto;
  position: relative;
}
img {
  display: block;
  margin: 15px 0;
  cursor: pointer;
}

</style>
