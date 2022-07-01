// Router

import { createWebHistory, createRouter } from "vue-router";   // HTML5 모드
import StartView from './components/StartView.vue';
import Login from './components/Login.vue';
import ErrorPage from './components/ErrorPage.vue';
import Register from './components/Register.vue';
import Main from './components/Main.vue';



const routes = [
  {
    path: "/",
    name: "StartView",
    component: StartView,
  },
  {
    path: "/login",    // 첫 화면을 로그인 화면으로 설정
    name: "Login",
    component: Login,
  },
  {
    path: "/register",  // 회원가입 페이지로 이동
    name: "Register",
    component: Register,
  },
  {
    path: "/main",
    name: "Main",
    component: Main,
  },
  // 같은 라우터에 걸리면 위에 있는걸 적용시킨다. 따라서, 404 페이지같은 경우 맨 아래에 기재하자.
  {
    path: "/:anyting",
    component: ErrorPage,
  },
];

const router = createRouter({
  history: createWebHistory(),   //^ HTML5 모드
  /*
  HTML5 모드인 경우,
  서버에 아무 기능을 개발안해놨으면 404페이지가 뜨거나 그럴 수 있다.  
  그래서 이걸 방지하려면 서버에다가 "어떤 사람이 /어쩌구로 접속하면 그냥 Vue에게 라우팅 맡겨주세요" 라고 미리 기능개발이 필요하다.
  */
  routes,
});

export default router; 