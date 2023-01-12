
const Path = require('path');
const Webpack = require('webpack');

const mode = 'development';
const minimize = (mode == 'development' ? false : true);

module.exports = [
	{
		mode: mode,
		target: 'web',
		devtool: false,
	
		entry: {
			'background': './extension/background.ts',
			'frontend': './extension/frontend.ts',
			'options': './extension/options.ts',
		},

		module: {
			rules: [
				{
					test: /\.tsx?$/,
					exclude: /node_modules/,
					loader: 'ts-loader',
					options: {
						projectReferences: true,
						configFile: 'tsconfig.json',
					}
				},
			],
		},

		resolve: {
			extensions: [ '.tsx', '.ts', '.jsx', '.js' ],

			//
			// Polyfill node modules for mqtt.js. References:
			// - https://webpack.js.org/configuration/resolve/#resolvefallback
			// - https://www.alchemy.com/blog/how-to-polyfill-node-core-modules-in-webpack-5
			// - https://viglucci.io/how-to-polyfill-buffer-with-webpack-5
			//

			fallback: {
				buffer: require.resolve('buffer/'),
				process: require.resolve('process/browser'),
				url: require.resolve("url/"),
			},
		},

		optimization: {
			minimize: minimize,
			concatenateModules: true,
		},

		output: {
			path: Path.join(Path.resolve(__dirname), 'build'),
			filename: '[name].js',
		},

	  plugins: [
			//
			// Polyfill node modules for mqtt.js. References:
			// - https://webpack.js.org/configuration/resolve/#resolvefallback
			// - https://www.alchemy.com/blog/how-to-polyfill-node-core-modules-in-webpack-5
			// - https://viglucci.io/how-to-polyfill-buffer-with-webpack-5
			//

			new Webpack.ProvidePlugin({
				Buffer: ['buffer', 'Buffer'],
				process: 'process/browser',
			}),
  	],
	},

	{
		mode: mode,
		target: 'web',
		devtool: false,
	
		entry: {
			'query': { import: './extension/content_scripts/query.ts', library: { type: 'assign', name: 'Query' } },
			'wait':  { import: './extension/content_scripts/wait.ts',  library: { type: 'assign', name: 'Wait' } },
		},

		module: {
			rules: [
				{
					test: /\.tsx?$/,
					exclude: /node_modules/,
					loader: 'ts-loader',
					options: {
						projectReferences: true,
						configFile: 'tsconfig.json',
					}
				},
			],
		},

		resolve: {
			extensions: [ '.tsx', '.ts', '.jsx', '.js' ],
		},

		optimization: {
			minimize: minimize,
			concatenateModules: true,
		},

		output: {
			path: Path.join(Path.resolve(__dirname), 'build', 'content_scripts'),
			filename: '[name].js',
			library: {
				type: 'var',
			}
		},
	}
]

