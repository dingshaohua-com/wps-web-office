import { axiosInstance } from '@/api/custom-axios';
import { useEffect, useState } from 'react';
import { Link } from 'react-router';

type FileItem = {
  id: string;
  name: string;
  size: number;
};

type ApiResponse<T> = {
  code: number;
  data: T;
};

function formatFileSize(size: number): string {
  return `${(size / 1024).toFixed(2)} KB`;
}

export default function Home() {
  const [files, setFiles] = useState<FileItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  const loadFiles = async () => {
    try {
      const response = await axiosInstance.get<ApiResponse<FileItem[]>>('/test-office');
      setFiles(response.data.data)
    } catch {
      setError('文件列表加载失败')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadFiles()
  }, []);

  return (
    <main className="min-h-screen bg-slate-50 px-6 py-12 text-slate-900">
      <section className="mx-auto max-w-2xl">
        <h1 className="mb-6 text-2xl font-semibold">文件列表</h1>

        <div className="overflow-hidden rounded-xl border border-slate-200 bg-white shadow-sm">
          {loading && <p className="p-6 text-sm text-slate-500">正在加载...</p>}
          {error && <p className="p-6 text-sm text-red-600">{error}</p>}
          {!loading && !error && files.length === 0 && (
            <p className="p-6 text-sm text-slate-500">暂无文件</p>
          )}

          {!loading && !error && files.length > 0 && (
            <ul className="divide-y divide-slate-100">
              {files.map((file) => (
                <li key={file.id}>
                  <Link
                    className="flex items-center justify-between px-6 py-4 transition hover:bg-slate-50"
                    to={`/preview/${encodeURIComponent(file.id)}?name=${encodeURIComponent(file.name)}`}
                  >
                    <span className="font-medium">{file.name}</span>
                    <span className="text-sm text-slate-500">{formatFileSize(file.size)}</span>
                  </Link>
                </li>
              ))}
            </ul>
          )}
        </div>
      </section>
    </main>
  );
}
