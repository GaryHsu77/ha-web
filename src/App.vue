<template>
  <div>
    <!-- title  -->
    <div id="app">
      <img alt="logo" src="./assets/logo.png" height="600px" width="1000px">
    </div>
    <div>
      <!-- devices table  -->
      <div class="content" :style="{ 'background-color': '#000000' }">
        <div class="md-layout">
          <div class="md-layout-item md-medium-size-100 md-xsmall-size-100 md-size-100">
            <md-card>
              <md-card-header id="header">
                <h3 class="title">Device List</h3>
              </md-card-header>
              <md-card-content>
                <md-table v-model="devices" @md-selected="onSelect">
                  <md-table-row slot="md-table-row" slot-scope="{ item }" :class="getClass(item)" md-selectable="single">
                    <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
                    <md-table-cell md-label="Model">{{ item.model }}</md-table-cell>
                    <md-table-cell md-label="UUID">{{ item.uuid }}</md-table-cell>
                    <md-table-cell md-label="Status">{{ item.status }}</md-table-cell>
                  </md-table-row>
                </md-table>
              </md-card-content>
            </md-card>
          </div>
        </div>
      </div>
      <!-- device table -->
      <div class="content" v-if="isSelected">
        <div class="md-layout">
          <!-- device card -->
          <div class="md-layout-item md-medium-size-100 md-size-33">
            <md-card class="md-card-profile">
              <div class="md-card-avatar">
                <img class="img" :src="cardUserImage" />
              </div>
              <md-card-content>
                <h3 class="category text-gray">{{ selected.model }}</h3>
                <h4 class="card-title">{{ selected.name }}</h4>
                <p class="card-description">
                  IP: {{ selected.deviceIp }}
                </p>
                <p class="card-description">
                  UUID: {{ selected.uuid }}
                </p>
                <md-button class="md-raised md-success" v-if="selected.statusb">master</md-button>
                <md-button class="md-raised md-accent" v-if="!selected.statusb">master</md-button>
                <md-button class="md-raised md-accent" v-if="selected.statusb">backup</md-button>
                <md-button class="md-raised md-danger" v-if="!selected.statusb">backup</md-button>
              </md-card-content>
            </md-card>
          </div>
          <!-- services table -->
          <div class="md-layout-item md-medium-size-100 md-size-66">
            <md-card>
              <md-card-header id="header">
                <h4 class="title">Service List</h4>
              </md-card-header>
              <md-card-content>
                <md-table v-model="services" @md-selected="onSelect">
                  <md-table-row slot="md-table-row" slot-scope="{ item }" :class="getClass(item)" md-selectable="single">
                    <md-table-cell md-label="Backup"><md-switch v-model="selected.statusb" class="md-danger"></md-switch></md-table-cell>
                    <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
                    <md-table-cell md-label="Model">{{ item.version }}</md-table-cell>
                    <md-table-cell md-label="Status">{{ item.state }}</md-table-cell>
                    <md-table-cell md-label="Configuration Sync Time">{{ item.syncTime }}</md-table-cell>
                  </md-table-row>
                </md-table>
              </md-card-content>
            </md-card>
          </div>
        </div>
        <h4>{{ selected }}</h4>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'app',
  components: {},
  props: {},
  data: () => ({
    isSelected: false,
    selected: {},
    devices: [],
    services: [],
    cardUserImage: '',
    timer: {}
  }),
  methods: {

    getClass: ({ id }) => ({
      'md-info': id === 0
    }),

    onSelect (item) {
      if (item === undefined) {
        return
      }
      this.selected = item
      this.services = item['services']
      this.isSelected = true
      this.cardUserImage = require('@/assets/img/' + item.model + '.png')
    },

    timerStart () {
      return setInterval(() => {
        axios.get('/slaves').then(response => {
          this.devices = []
          var devs = JSON.parse(JSON.stringify(response.data))
          var i = 0
          for (var key in devs) {
            var dev = {
              id: i,
              name: devs[key]['deviceName'],
              model: devs[key]['modelName'],
              uuid: devs[key]['deviceId'],
              status: 'connected',
              statusb: true,
              services: []
            }
            var si = 0
            for (var svc in devs[key]['dpgrs']) {
              var s = {
                id: si,
                name: devs[key]['dpgrs'][svc]['name'],
                version: '2.0.0',
                state: 'good',
                syncTime: devs[key]['dpgrs'][svc]['lastUpdated']
              }
              switch (devs[key]['dpgrs'][svc]['status']) {
                case 1:
                  s['state'] = 'online'
                  break
                case 2:
                  s['state'] = 'offline'
                  dev['status'] = 'disconnect'
                  dev['statusb'] = false
                  break
                case 3:
                  s['state'] = 'updating'
                  break
                default:
                  console.log('unknown state')
              }
              dev['services'].push(s)
              dev['deviceIp'] = devs[key]['dpgrs'][svc]['ipaddr']
              si++
            }
            this.devices.push(dev)
            i++
          }
        })
      }, 1000)
    }
  },
  created: function () {
    console.log('created')
    this.timer = this.timerStart()
  },
  destroyed () {
    clearInterval(this.timer)
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
#header{
    background-color:#16a1a1;
}
.md-switch {
  display: flex;
}
</style>
