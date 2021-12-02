import axios from 'axios';
import { serializeCalendar } from '../../functions/serializers';
import qs from 'qs';

const apiUrl = 'https://api.shimo0108-app.com:9999';
const formHeader = { headers: { 'content-type': 'application/x-www-form-urlencoded' } };


const state = {
  calendars: [],
  calendar: null,
};

const getters = {
  calendars: (state) => state.calendars.map((calendar) => serializeCalendar(calendar)),
  calendar: (state) => serializeCalendar(state.calendar),
};

const mutations = {
  setCalendars: (state, calendars) => (state.calendars = calendars),
  appendCalendar: (state, calendar) => (state.calendars = [...state.calendars, calendar]),
  updateCalendar: (state, calendar) => (state.calendars = state.calendars.map((c) => (c.id === calendar.id ? calendar : c))),
  removeCalendar: (state, calendar) => (state.calendars = state.calendars.filter((c) => c.id !== calendar.id)),
  setCalendar: (state, calendar) => (state.calendar = calendar),
};

const actions = {
  async fetchCalendars({ commit }) {
    const response = await axios.get(`${apiUrl}/api/v1/calendars`);
    commit('setCalendars', response.data);
  },
  async createCalendar({ commit }, calendar) {
    console.log(calendar)
    const response = await axios.post(`${apiUrl}/api/v1/calendars`, qs.stringify(calendar), formHeader);
    commit('appendCalendar', response.data);
  },
  async updateCalendar({ commit }, calendar) {
    const response = await axios.put(`${apiUrl}/api/v1/calendars/${calendar.id}`, qs.stringify(calendar), formHeader);
    commit('updateCalendar', response.data);
  },
  async deleteCalendar({ commit }, id) {
    console.log(id)
    const response = await axios.delete(`${apiUrl}/api/v1/calendars/${id}`);
    commit('removeCalendar', response.data);
  },
  setCalendar({ commit }, calendar) {
    commit('setCalendar', calendar);
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
