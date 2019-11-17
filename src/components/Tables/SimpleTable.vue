<template>
  <div>
    <md-table v-model="devices" :table-header-color="tableHeaderColor" @md-selected="onSelect">
      <md-table-row slot="md-table-row" slot-scope="{ item }" :class="getClass(item)" md-selectable="single">
        <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
        <md-table-cell md-label="Model">{{ item.model }}</md-table-cell>
        <md-table-cell md-label="UUID">{{ item.uuid }}</md-table-cell>
        <md-table-cell md-label="Status">{{ item.status }}</md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
export default {
  name: 'simple-table',
  props: {
    tableHeaderColor: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      selected: {},
      devices: []
    }
  },
  methods: {
    getClass: ({ id }) => ({
      'md-info': id === 0
    }),
    onSelect (item) {
      this.selected = item
      var response = {
        selected: this.selected,
        devices: this.devices
      }
      this.$emit('input', response)
    }
  }
}
</script>
