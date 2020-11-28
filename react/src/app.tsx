import React from 'react';
import { BasicLayoutProps, Settings as LayoutSettings, PageLoading } from '@ant-design/pro-layout';
import Footer from '@/components/Footer';
import defaultSettings from '../config/defaultSettings';

export const initialStateConfig = {
  loading: <PageLoading />,
};

export async function getInitialState(): Promise<{
  settings?: LayoutSettings;
}> {
  return {
    settings: defaultSettings,
  };
}

export const layout = ({
  initialState,
}: {
  initialState: { settings?: LayoutSettings };
}): BasicLayoutProps => {
  return {
    rightContentRender: () => <></>,
    disableContentMargin: false,
    footerRender: () => <Footer />,
    onPageChange: () => {},
    menuHeaderRender: undefined,
    ...initialState?.settings,
  };
};
