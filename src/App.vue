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
                <sweet-modal icon="success" hide-close-button overlay-theme="dark" modal-theme="dark" ref="replaceSuccess">
                  Success!
                </sweet-modal>
                <sweet-modal icon="error" hide-close-button overlay-theme="dark" modal-theme="dark" ref="replaceError">
                  This is an error…
                </sweet-modal>
                <md-table v-model="devices" @md-selected="onSelect">
                  <md-table-row slot="md-table-row" slot-scope="{ item }" :class="getClass(item)" md-selectable="single">
                    <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
                    <md-table-cell md-label="Status" class="text-success" v-if="item.statusb">{{ item.status }}</md-table-cell>
                    <md-table-cell md-label="Status" class="text-danger" v-if="!item.statusb">{{ item.status }}</md-table-cell>
                    <md-table-cell md-label="IP">{{ item.deviceIp }}</md-table-cell>
                    <md-table-cell md-label="Model">{{ item.model }}</md-table-cell>
                    <md-table-cell md-label="UUID">{{ item.uuid }}</md-table-cell>
                    <md-table-cell md-label="Setting" v-if="replacePress">
                      <md-button class="md-sm md-info" @click="replaceClose(item)" v-if="replacePress&&(item.uuid!==selected.uuid)">TO BE POSSESSED</md-button>
                    </md-table-cell>
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
                <div>
                <md-button class="md-sm md-info" @click="onReplace()" v-if="!replacePress">POSSESS</md-button>
                </div>
                <md-button class="md-raised md-success" v-if="selected.statusb">master</md-button>
                <md-button class="md-raised md-accent" v-if="!selected.statusb">master</md-button>
                <md-button class="md-raised md-accent" v-if="selected.statusb">backup</md-button>
                <md-button class="md-raised md-warning" v-if="!selected.statusb">backup</md-button>
              </md-card-content>
            </md-card>
          </div>
          <div class="md-layout-item md-medium-size-100 md-size-66">
            <!-- services table -->
            <md-card>
              <md-card-header id="header">
                <h4 class="title">Service List</h4>
              </md-card-header>
              <md-card-content>
                <md-table v-model="services" @md-selected="serviceOnSelect">
                  <md-table-row slot="md-table-row" slot-scope="{ item }" :class="getClass(item)" md-selectable="single">
                    <md-table-cell md-label="Backup">
                      <my-switch v-model="item.enable" open-name="" close-name="" color="green" @change="serviceOnSelect"></my-switch>
                    </md-table-cell>
                    <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
                    <md-table-cell md-label="Model">{{ item.version }}</md-table-cell>
                    <md-table-cell md-label="Status" class="text-success" v-if="item.stateb">{{ item.state }}</md-table-cell>
                    <md-table-cell md-label="Status" class="text-danger" v-if="!item.stateb">{{ item.state }}</md-table-cell>
                    <md-table-cell md-label="Configuration Sync Time">{{ item.syncTime }}</md-table-cell>
                  </md-table-row>
                </md-table>
              </md-card-content>
            </md-card>
            <!-- configuration area -->
            <md-card v-if="showConfig">
              <md-card-header id="header">
                <h4 class="title">Configuration</h4>
              </md-card-header>
              <md-card-content>
                <md-field>
                  <md-textarea v-model="configuration"></md-textarea>
                </md-field>
                <sweet-modal icon="success" hide-close-button overlay-theme="dark" modal-theme="dark" ref="saveSuccess">
                  Success!
                </sweet-modal>
                <sweet-modal icon="error" hide-close-button overlay-theme="dark" modal-theme="dark" ref="saveError">
                  This is an error…
                </sweet-modal>
                <md-button class="md-raised md-sm md-info" @click="onSave()">SAVE</md-button>
              </md-card-content>
            </md-card>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { SweetModal } from 'sweet-modal-vue'
import mySwitch from 'vue-switch/switch-2.vue'
import axios from 'axios'

export default {
  name: 'app',
  components: {
    'my-switch': mySwitch,
    SweetModal
  },
  props: {},
  data: () => ({
    showConfig: false,
    configurationUUID: '',
    configuration: {},
    replacePress: false,
    isSelected: false,
    selected: {},
    selectedService: {},
    original: '',
    data: {},
    devices: [],
    services: [],
    cardUserImage: '',
    timer: {}
  }),
  methods: {
    replaceClose: function (item) {
      if (item === undefined) {
        return
      }
      var content = {}
      content['deviceID'] = item.uuid
      axios.put('/replace/' + this.selected.uuid, content).then(response => {
        if (response.status === 200) {
          this.$refs.replaceSuccess.open()
        } else {
          this.$refs.replaceError.open()
        }
      })
      this.replacePress = false
    },

    onReplace: function () {
      this.isSelected = false
      this.replacePress = true
    },

    onSave: function () {
      var content = {
        id: this.selectedService['id'],
        deviceid: this.selected['uuid'],
        data: this.configuration
      }
      axios.put('/apply', content).then(response => {
        if (response.status === 200) {
          this.$refs.saveSuccess.open()
        } else {
          this.$refs.saveError.open()
        }
      })
    },

    getClass: ({ id }) => ({
      'md-info': id === 0
    }),

    onSelect (item) {
      if (item === undefined) {
        return
      }
      if (item === null) {
        return
      }
      if (this.replacePress) {
        return
      }
      this.selected = item
      this.services = item['services']
      this.isSelected = true
      this.cardUserImage = require('@/assets/img/' + item.model + '.png')
      this.showConfig = false
    },

    serviceOnSelect (item) {
      if (item === undefined) {
        return
      }
      for (var i = 0; i < this.devices.length; i++) {
        if (this.devices[i]['uuid'] === this.selected['uuid']) {
          this.devices[i] = this.selected
          break
        }
      }
      if (this.original !== JSON.stringify(this.devices)) {
      }
      this.showConfig = true
      this.configuration = JSON.stringify(JSON.parse(item.config), undefined, 4)
      this.selectedService = item
    },

    timerStart () {
      return setInterval(() => {
        axios.get('/slaves').then(response => {
          this.data = response.data
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
            for (var svc in devs[key]['dpgrs']) {
              var s = {
                id: devs[key]['dpgrs'][svc]['id'],
                name: devs[key]['dpgrs'][svc]['name'],
                version: '2.0.0',
                state: 'good',
                stateb: true,
                enable: devs[key]['dpgrs'][svc]['enable'],
                syncTime: devs[key]['dpgrs'][svc]['lastUpdated'],
                config: devs[key]['dpgrs'][svc]['config']
              }
              switch (devs[key]['dpgrs'][svc]['status']) {
                case 1:
                  s['state'] = 'online'
                  break
                case 2:
                  s['state'] = 'offline'
                  s['stateb'] = false
                  dev['status'] = 'disconnect'
                  dev['statusb'] = false
                  break
                case 3:
                  s['state'] = 'updating'
                  break
                default:
              }
              dev['services'].push(s)
            }
            dev['deviceIp'] = devs[key]['deviceIp']
            if (this.isSelected) {
              if (dev['uuid'] === this.selected['uuid']) {
                this.selected = dev
                this.services = dev['services']
              }
            }
            this.devices.push(dev)
            i++
          }
        })
        this.original = JSON.stringify(this.devices)
      }, 1000)
    }
  },
  created: function () {
    this.timer = this.timerStart()
  },
  destroyed () {
    clearInterval(this.timer)
  },
  watch: {
    selected: function () {
    }
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
