import axios from 'axios';
import { format } from 'date-fns';
import qs from 'qs';

const apiUrl = process.env.VUE_APP_API_BASE_URL;
const formHeader = { headers: { 'content-type': 'application/x-www-form-urlencoded' } };

const state = {
  events: [],
  event: null,
  isEditMode: false,
};

const getters = {
  events: (state) =>
    state.events.map((event) => {
      return {
        ...event,
        start: new Date(event.started_at.replace(/Z$/, '+09:00')),
        end: new Date(event.ended_at.replace(/Z$/, '+09:00')),
      };
    }),
  event: (state) =>
    state.event
      ? {
          ...state.event,
          start: new Date(state.event.started_at),
          end: new Date(state.event.ended_at),
          startDate: format(new Date(state.event.start), 'yyyy-MM-dd'),
          endDate: format(new Date(state.event.end), 'yyyy-MM-dd'),
          startTime: format(new Date(state.event.start), 'HH:mm'),
          endTime: format(new Date(state.event.end), 'HH:mm'),
          color: event.color || '#2196F3',
        }
      : null,
  isEditMode: (state) => state.isEditMode,
};

const mutations = {
  setEvents: (state, events) => (state.events = events),
  removeEvent: (state, event) => (state.events = state.events.filter((e) => e.id !== event.id)),
  resetEvent: (state) => (state.event = null),
  updateEvent: (state, event) => (state.events = state.events.map((e) => (e.id === event.id ? event : e))),
  appendEvent: (state, event) => (state.events = [...state.events, event]),

  setEvent: (state, event) => (state.event = event),
  setEditMode: (state, bool) => (state.isEditMode = bool),
};

const actions = {
  async fetchEvents({ commit }) {
    console.log(apiUrl)
    const response = await axios.get(`${apiUrl}/events`);

    commit('setEvents', response.data); // mutationを呼び出す
  },
  async createEvent({ commit }, event) {
    event.start = format(new Date(event.start.toString() + ':00'), 'yyyy-MM-dd HH:mm:00');
    event.end = format(new Date(event.end.toString() + ':00'), 'yyyy-MM-dd HH:mm:00');
    const response = await axios.post(`${apiUrl}/events`, qs.stringify(event), formHeader);
    response.data.started_at = response.data.started_at.replace(/Z$/, '+09:00');
    response.data.ended_at = response.data.ended_at.replace(/Z$/, '+09:00');
    console.log(response.data);
    commit('appendEvent', response.data);
  },
  async updateEvent({ commit }, event) {
    event.start = format(new Date(event.start.toString() + ':00'), 'yyyy-MM-dd HH:mm:00');
    event.end = format(new Date(event.end.toString() + ':00'), 'yyyy-MM-dd HH:mm:00');
    const response = await axios.put(`${apiUrl}/events/${event.id}`, qs.stringify(event), formHeader);
    response.data.started_at = response.data.started_at.replace(/Z$/, '+09:00');
    response.data.ended_at = response.data.ended_at.replace(/Z$/, '+09:00');
    commit('updateEvent', response.data);
  },
  async deleteEvent({ commit }, id) {
    const response = await axios.delete(`${apiUrl}/events/${id}`);
    commit('removeEvent', response.date);
    commit('resetEvent');
  },
  setEvent({ commit }, event) {
    commit('setEvent', event);
  },
  setEditMode({ commit }, bool) {
    commit('setEditMode', bool);
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
