import { lazy } from 'react';
import { createHashRouter } from 'react-router';
import Root from '@/components/root';

const router = createHashRouter([
  {
    path: '/',
    Component: Root,
    children: [
      { index: true, Component: lazy(() => import('@/pages/home')) },
      { path: '/preview/:fileId', Component: lazy(() => import('@/pages/preview')) },
    ],
  },
]);

export default router;
