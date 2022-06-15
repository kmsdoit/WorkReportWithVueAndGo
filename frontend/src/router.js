// Router

import { createWebHistory, createRouter } from "vue-router";   // HTML5 모드
// import App from './App.vue';
import Login from './components/Login.vue';
import ErrorPage from './components/ErrorPage.vue';


const routes = [
  // {
  //   path: "/",
  //   component: App,
  // },
  {
    path: "/login",
    name: "Login",
    component: Login,
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