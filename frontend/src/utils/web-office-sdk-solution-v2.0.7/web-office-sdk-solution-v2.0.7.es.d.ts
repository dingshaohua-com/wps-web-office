import type {
  IAppConfig,
  IConfig,
  IWps,
  OfficeType,
} from './index';

type WebOfficeInitConfig = Omit<IAppConfig, 'attrAllow'> & {
  attrAllow?: string | string[];
};

interface WebOfficeSDKStatic {
  readonly version: string;
  readonly OfficeType: OfficeType;
  init(config: WebOfficeInitConfig): IWps;
  config(config: IConfig): IWps | undefined;
}

declare const WebOfficeSDK: WebOfficeSDKStatic;

export default WebOfficeSDK;
export * from './index';
