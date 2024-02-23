/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["internal/templates/**/*.templ"],
  plugins: [require("daisyui"), require("@tailwindcss/forms")],
};
