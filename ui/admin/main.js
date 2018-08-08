
window.appUrl = '/admin'

import Vue from 'vue';
import iView from 'iview';
import {router} from './libs/router/index';
import App from './libs/frame/app.vue';
import axios from "axios";
import Cookies from 'js-cookie';
import 'iview/dist/styles/iview.css';

Vue.use(iView);

Vue.prototype.ajax = function (url, p, fun) {
    let th = this;
    var pp = p ? p : {}
    pp["siteId"] = window.gk.siteId;
    axios.post('/api' + url, pp).then(function (r) {
        for (var k in r.data) {
            if (th.hasOwnProperty(k))
                th[k] = r.data[k]
        }
        if (th.hasOwnProperty("loading")) {
            th.loading = false;
        }
        if (fun && typeof(fun) == 'function') {
            fun(r.data, th)
        }
        vm.$emit("data", window.gk)
    }).catch((err) => {
        if (err.response && err.response.status == 401) {
            window.goUrl = window.location.hash
            window.gk.login = false;
            Cookies.remove('webGeekAdmin');
            th.$router.replace(appUrl + '/p/login')
            vm.$emit("data", window.gk)
        } else {
            console.log(err)
        }
    })
}

window.goUrl = '/';
window.gk = {login: false, siteId: 0};

Vue.component('go', function (resolve) {
    require(['./libs/go.vue'], resolve)
})
window.vm = new Vue({
    el: '#app', router, render: h => h(App),
});
