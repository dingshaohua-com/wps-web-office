import WebOfficeSDK from '@/utils/web-office-sdk-solution-v2.0.7/web-office-sdk-solution-v2.0.7.es.js';
import { ArrowLeft } from 'lucide-react';
import { useEffect, useRef, useState } from 'react';
import { Link, useParams, useSearchParams } from 'react-router';

const APP_ID = import.meta.env.VITE_WPS_APP_ID || 'SX20261009BIXNZA';

export default function Preview() {
  const { fileId } = useParams();
  const [searchParams] = useSearchParams();
  const fileName = searchParams.get('name') || fileId || '文件预览';
  const containerRef = useRef<HTMLDivElement>(null);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!containerRef.current || !fileId) return;

    const officeType = fileName.toLowerCase().endsWith('.pdf')
      ? WebOfficeSDK.OfficeType.Pdf
      : WebOfficeSDK.OfficeType.Writer;
    const instance = WebOfficeSDK.init({
      officeType,
      appId: APP_ID,
      fileId,
      mount: containerRef.current,
    });

    instance.on('error', (event: unknown) => {
      console.error('WebOffice error:', event);
      setError('文件预览失败');
    });

    return () => {
      void instance.destroy();
    };
  }, [fileId, fileName]);

  if (!fileId) {
    return <p className="p-6 text-sm text-red-600">缺少文件 ID</p>;
  }

  return (
    <main className="flex min-h-screen flex-col bg-slate-100">
      <header className="flex h-14 items-center gap-4 border-b border-slate-200 bg-white px-5">
        <Link
          aria-label="返回文件列表"
          className="rounded-md p-2 text-slate-500 transition hover:bg-slate-100 hover:text-slate-900"
          title="返回文件列表"
          to="/"
        >
          <ArrowLeft aria-hidden="true" size={20} />
        </Link>
        <h1 className="truncate font-medium text-slate-900">{fileName}</h1>
      </header>

      {error && <p className="bg-red-50 px-5 py-2 text-sm text-red-600">{error}</p>}
      <div className="flex flex-1 justify-center p-4">
        <div ref={containerRef} className="min-h-[calc(100vh-5.5rem)] w-full max-w-6xl bg-white" />
      </div>
    </main>
  );
}
