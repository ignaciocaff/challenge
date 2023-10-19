/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,ts}",
  ],
  theme: {
    extend: {
      colors: {
        appPrimary: "#38BBEC",
        appSecondary: "#3861EC",
        appAccent: "#257B9B",
        appNeutral: "#8C8CA1",
        appBase100: "#FAFCFE",
        appBase200: "#ECF1F4",
        appNeutralFocus: "#ECF1F4",
        appInfo: "#44E2FF",
        appSuccess: "#2ECC00",
        appWarning: "#EC6938",
        appError: "#EC3861",
      },
      fontFamily: {
        sans: ['GothamRounded', 'sans-serif'],
      },
    },
  },
  daisyui: {
    themes: [
      {
        mytheme: {
          "primary": "#38BBEC",
          "secondary": "#3861EC",
          "accent": "#257B9B",
          "neutral": "#8C8CA1",
          "base-100": "#FAFCFE",
          "base-200": "#ECF1F4",
          "neutral-focus": "#ECF1F4",
          "info": "#44E2FF",
          "success": "#2ECC00",
          "warning": "#EC6938",
          "error": "#EC3861",
        },
      },
    ],
  },
  plugins: [require("daisyui")],
}
