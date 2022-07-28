<template>
  <div class="work-container">
    <div class="inner">
      
      <div class="start-work">
        <h3>name님,</h3>
        <div v-if="workStep == 0">
          <!-- <p>{{출근.출근전}}</p> -->
        </div>
        <div v-if="workStep == 1">
          <!-- <p>{{출근.출근후}}</p> -->
        </div>
      </div>

      <div class="end-work">
        <div class="working">
          <p>업무중 : </p>
          <div>
            <p>출근시간 : </p>
            <p>퇴근시간 : </p>
          </div>
          
        </div>
        <div class="btn-wraper">
          <button class="">자리비움</button>
          <button class="">업무종료</button>
        </div>
      </div>

    </div>
  </div>
</template>
<script>
export default {
  name: 'Work',
  data() {
    return {
      // 나중에 영어로 바꾸자

      // 출근: {
      //   출근전: '아직 출근 전입니다 :)',
      //   출근후: '출근 완료!'
      // },
      // workStep: 0,
    }
  },
  
  // mounted() {
    // 로그인 했는지 안했는지 검사해야 함!!
    // 로그인 한 경우에만 보여주기 위해
    // 로그인 했으면 보여주고, 안했으면 로그인창으로 이동
  // },
}
</script>
<style>
  .work-container {
    width: 83vw;
    height: 100vh;
    /* background: #333; */
    right: 0;
  }
  .inner {
    width: 80%;
    height: 100%;
    display: block;
    margin: 0 auto;
    /* background: #eee; */
    position: relative;
  }


  .start-work {
    width: 45%;
    height: 300px;
    /* background: #000; */
    position: absolute;
    left: 40px;
    top: 50px;
    /* border: 1px solid #333; */
    border: none;
    border-radius: 5px;
    box-shadow: 2px 2px 5px lightgray;
  }



  .end-work {
    width: 45%;
    height: 300px;
    /* background: #000; */
    position: absolute;
    right: 40px;
    top: 50px;
    border: none;
    border-radius: 5px;
    box-shadow: 2px 2px 5px lightgray;
  }
  .working {
    width: 90%;
    height: 150px;
    background-color: #b3e5fb;
    display: block;
    margin: 30px auto 0;
    border-radius: 5px;
  }
  .btn-wraper {
    width: 90%;
    height: 80px;
    background-color: #b3e5fb;
    display: block;
    margin: 15px auto 0;
    border-radius: 5px;
  }
</style>