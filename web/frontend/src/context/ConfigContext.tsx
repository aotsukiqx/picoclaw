import React, { createContext, useContext, useState, useEffect, useCallback } from 'react';

export interface AppConfig {
  voice: {
    enabled: boolean;
    asrServer: string;
    asrModel: string;
  };
  vision: {
    enabled: boolean;
    captureMode: 'photo' | 'video_3s';
  };
  gemma4ApiKey: string;
  gemma4Model: string;
}

const defaultConfig: AppConfig = {
  voice: {
    enabled: true,
    asrServer: 'ws://localhost:10388',
    asrModel: 'whisper-v3',
  },
  vision: {
    enabled: true,
    captureMode: 'photo',
  },
  gemma4ApiKey: '',
  gemma4Model: 'google/gemma-4-E4B-it',
};

interface ConfigContextValue {
  config: AppConfig;
  updateConfig: (patch: Partial<AppConfig>) => Promise<void>;
  isLoading: boolean;
}

const ConfigContext = createContext<ConfigContextValue | null>(null);

export function ConfigProvider({ children }: { children: React.ReactNode }) {
  const [config, setConfig] = useState<AppConfig>(defaultConfig);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const stored = localStorage.getItem('@picoclaw_config');
    if (stored) {
      try {
        setConfig({ ...defaultConfig, ...JSON.parse(stored) });
      } catch {
        setConfig(defaultConfig);
      }
    }
    setIsLoading(false);
  }, []);

  const updateConfig = useCallback(async (patch: Partial<AppConfig>) => {
    const updated = { ...config, ...patch };
    await localStorage.setItem('@picoclaw_config', JSON.stringify(updated));
    setConfig(updated);
  }, [config]);

  return (
    <ConfigContext.Provider value={{ config, updateConfig, isLoading }}>
      {children}
    </ConfigContext.Provider>
  );
}

export function useConfig() {
  const context = useContext(ConfigContext);
  if (!context) {
    throw new Error('useConfig must be used within ConfigProvider');
  }
  return context;
}
