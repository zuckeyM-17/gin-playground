'use strict'

const csfrToken = document.querySelector('meta[name="csrf-token"]').getAttribute('content')

const app = new Vue({
  el: "#app",
  delimiters: ['[[', ']]'],
  data: {
    message: "hello Vue.js"
  },
  methods: {
    signIn() {
      axios.post('/api/sign_in', {},
        { headers: { 'X-CSRF-TOKEN': csfrToken }, }
      )
    }
  }
})