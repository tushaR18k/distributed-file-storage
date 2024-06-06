// app/pages/upload.tsx

import Head from 'next/head'
import Link from 'next/link'

export default function Upload() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center py-2">
      <Head>
        <title>Upload Files</title>
        <meta name="description" content="Upload your files to the distributed storage system" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex w-full flex-1 flex-col items-center justify-center px-20 text-center">
        <h1 className="text-6xl font-bold">Upload Files</h1>
        <p className="mt-3 text-2xl">Use the form below to upload your files.</p>

        <form className="mt-6 flex w-full max-w-4xl flex-col items-center justify-around sm:w-full">
          <input type="file" multiple className="block w-full text-sm text-slate-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-violet-50 file:text-violet-700 hover:file:bg-violet-100"/>
          <Link href="/files" className="mt-6 w-32 rounded-xl bg-blue-600 p-2 text-white hover:bg-blue-700">
              Upload
          </Link>
        </form>
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
