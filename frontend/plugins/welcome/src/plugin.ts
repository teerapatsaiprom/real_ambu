import { createPlugin } from '@backstage/core';
import Login from './components/Login';
import CreatePredicament from './components/Createpredict';
import ResultPage from './components/Result';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/login', Login);
    router.registerRoute('/user', CreatePredicament);
    router.registerRoute('/results', ResultPage);
  },
});
