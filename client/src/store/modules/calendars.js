import axios from 'axios';
import { serializeCalendar } from '../../functions/serializers';
import qs from 'qs';

const apiUrl = process.env.VUE_APP_API_BASE_URL;
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
    const response = await axios.get(`${apiUrl}/calendars`);
    commit('setCalendars', response.data);
  },
  async createCalendar({ commit }, calendar) {
    const response = await axios.post(`${apiUrl}/calendars`, qs.stringify(calendar), formHeader);
    commit('appendCalendar', response.data);
  },
  async updateCalendar({ commit }, calendar) {
    const response = await axios.put(`${apiUrl}/calendars/${calendar.id}`, qs.stringify(calendar), formHeader);
    commit('updateCalendar', response.data);
  },
  async deleteCalendar({ commit }, id) {
    console.log(id);
    const response = await axios.delete(`${apiUrl}/calendars/${id}`);
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
