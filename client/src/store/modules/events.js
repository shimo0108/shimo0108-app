import axios from 'axios';
import { isDateWithinInterval, compareDates } from '../../functions/datetime';
import { serializeEvent } from '../../functions/serializers';
import qs from 'qs';


const apiUrl = 'http://localhost:9999';
const formHeader = { headers: { 'content-type': 'application/x-www-form-urlencoded' } };


const state = {
  events: [],
  event: null,
  isEditMode: false,
  clickedDate: null,
};

const getters = {
  events: (state) => state.events.filter((event) => event.calendar).map((event) => serializeEvent(event)),
  event: (state) => serializeEvent(state.event),
  dayEvents: (state) =>
    state.events
      .map((event) => serializeEvent(event))
      .filter((event) => isDateWithinInterval(state.clickedDate, event.startDate, event.endDate))
      .sort(compareDates),
  isEditMode: (state) => state.isEditMode,
  clickedDate: (state) => state.clickedDate,
};

const mutations = {
  setEvents: (state, events) => (state.events = events),
  appendEvent: (state, event) => (state.events = [...state.events, event]),
  setEvent: (state, event) => (state.event = event),
  removeEvent: (state, event) => (state.events = state.events.filter((e) => e.id !== event.id)),
  resetEvent: (state) => (state.event = null),
  updateEvent: (state, event) => (state.events = state.events.map((e) => (e.id === event.id ? event : e))),
  setEditMode: (state, bool) => (state.isEditMode = bool),
  setClickedDate: (state, date) => (state.clickedDate = date),
};

const actions = {
  async fetchEvents({ commit }) {
    const response = await axios.get(`${apiUrl}/api/v1/events`);
    let responseEvents = response.data
    for (let event of responseEvents) {

      event.start = event.start_time.slice(0, -1) + '+09:00';
      event.end = event.end_time.slice(0, -1) + '+09:00';
      delete event.start_time;
      delete event.end_time;
    }
    commit('setEvents', responseEvents);
  },
  async createEvent({ commit }, event) {
    const response = await axios.post(`${apiUrl}/api/v1/events`, qs.stringify(event), formHeader);
    response.start = response.start_time;
    response.end = response.end_time;
    delete response.start_time;
    delete response.end_time;

    commit('appendEvent', event);
  },
  async deleteEvent({ commit }, id) {
    const response = await axios.delete(`${apiUrl}/api/v1/events/${id}`);
    commit('removeEvent', response.date);
    commit('resetEvent');
  },
  async updateEvent({ commit }, event) {
    const response = await axios.put(`${apiUrl}/api/v1/events/${event.id}`, qs.stringify(event), formHeader);
    commit('updateEvent', response.data);
  },
  setEvent({ commit }, event) {
    commit('setEvent', event);
  },
  setEditMode({ commit }, bool) {
    commit('setEditMode', bool);
  },
  setClickedDate({ commit }, date) {
    commit('setClickedDate', date);
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
