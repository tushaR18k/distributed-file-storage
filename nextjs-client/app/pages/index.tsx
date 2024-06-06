// app/pages/index.tsx

import Head from 'next/head'
import { useRouter } from 'next/router'

export default function Home() {
  const router = useRouter();

  return (
    <div className="flex min-h-screen flex-col items-center justify-center py-2">
      <Head>
        <title>Distributed File Storage System</title>
        <meta name="description" content="A distributed file storage system" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex w-full flex-1 flex-col items-center justify-center px-20 text-center">
        <h1 className="text-6xl font-bold">
          Welcome to the Distributed File Storage System
        </h1>

        <p className="mt-3 text-2xl">
          Upload and manage your files efficiently.
        </p>

        <div className="mt-6 flex max-w-4xl flex-wrap items-center justify-around sm:w-full">
          <button className="mt-6 w-96 rounded-xl border p-6 text-left hover:text-blue-600 focus:text-blue-600" onClick={() => router.push('/upload')}>
            <h3 className="text-2xl font-bold">Upload Files &rarr;</h3>
            <p className="mt-4 text-xl">
              Upload your files to the distributed storage system.
            </p>
          </button>

          <button className="mt-6 w-96 rounded-xl border p-6 text-left hover:text-blue-600 focus:text-blue-600" onClick={() => router.push('/files')}>
            <h3 className="text-2xl font-bold">Manage Files &rarr;</h3>
            <p className="mt-4 text-xl">
              View and manage your uploaded files.
            </p>
          </button>
        </div>
      </main>

      <footer className="flex h-24 w-full items-center justify-center border-t">
        <a
          className="flex items-center justify-center gap-2"
          href="https://github.com/"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by{' '}
          <img src="/vercel.svg" alt="Vercel Logo" className="h-4" />
        </a>
      </footer>
    </div>
  )
}
