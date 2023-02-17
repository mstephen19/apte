module.exports = {
    env: { browser: true, es2021: true },
    extends: ['turbo', 'prettier', 'eslint:recommended', 'plugin:@typescript-eslint/recommended'],
    plugins: ['@typescript-eslint', 'eslint-plugin-react'],
    rules: {
        'react/jsx-key': 'warn',
        indent: ['error', 4, { SwitchCase: 1 }],
        'linebreak-style': ['error', 'unix'],
        quotes: ['error', 'single', { avoidEscape: true }],
        semi: ['error', 'always'],
        'no-console': 'warn',
        'arrow-parens': 'error',
        'no-control-regex': 'off',
        '@typescript-eslint/no-explicit-any': 'off',
        '@typescript-eslint/no-non-null-assertion': 'off',
        '@typescript-eslint/consistent-type-imports': ['error', { prefer: 'type-imports' }],
    },
    parserOptions: { ecmaVersion: 'latest', sourceType: 'module' },
};
