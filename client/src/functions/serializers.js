export const serializeCalendar = (calendar) => {
  if (calendar === null) {
    return null;
  }
  return {
    ...calendar,
    color: calendar.color || '#2196F3',
  };
};
