/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
		"./extension/frontend.html",
		"./extension/options.html",
	],

  theme: {
		colors: {
      none: 'transparent',
      current: 'currentColor',
      white: '#ffffff',
			black: '#000000',

			gray10: '#f9f9fa', gray20: '#ededf0', gray30: '#d7d7db',
			gray40: '#b1b1b3', gray50: '#737373', gray60: '#4a4a4f;',
			gray70: '#38383d', gray80: '#2a2a2e', gray90: '#0c0c0d',

			blue40: '#45a1ff', blue50: '#0a84ff', blue60: '#0060df',
			blue70: '#003eaa', blue80: '#002275', blue90: '#000f40',

			blue40t: '#45a1ff4c', blue50t: '#0a84ff4c', blue60t: '#0060df4c',
			blue70t: '#003eaa4c', blue80t: '#0022754c', blue90t: '#000f404c',

			red50: '#ff0039', red60: '#d70022', red70: '#a4000f',
			red80: '#5a0002', red90: '#3e0200',

			red50t: '#ff00394c', red60t: '#d700224c', red70t: '#a4000f4c',
			red80t: '#5a00024c', red90t: '#3e02004c',
		},

		fontFamily: {
			sans: [ 'sans' ],
			serif: [ 'sans-serif' ],
			mono: [ 'monospace' ],
		},

		spacing: {
			'0': '0',
			'inset': '0.8rem',
			'frame': '3px',
			'delta': 'calc(0.8rem - 3px)',
			'line': '1.2em',
			'halfline': '0.6em',
			'control': '6px',
		},

		width: {
			'0': '0',
			'full': '100%',
			'action': '24px',
			'timestamp': '6em',
		},

		height: {
			'0': '0',
			'full': '100%',
			'header': '2.8rem',
			'footer': '2.5rem',
			'control': '2.5em',
			'action': '24px',
		},

		minWidth: {
			'0': '0',
			'button': '132px',
		},

		borderWidth: {
			'0': '0',
			'input': '1px',
			'focus': '2px',
			'danger': '2px',
		},

		outlineWidth: {
			'0': '0',
			'focus': '3px',
			'danger': '3px',
		},

		gap: {
			'0': '0',
			'line': '1.2em',
			'control': '6px',
		},

		extend: {
			lineHeight: {
				'header': '2.8rem',
				'footer': '2.5rem',
			},
		}
  },

  plugins: [],
}
