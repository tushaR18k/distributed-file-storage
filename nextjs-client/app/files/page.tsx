// app/pages/files.tsx

import Head from 'next/head'
import Link from 'next/link'

export default function Files() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center py-2">
      <Head>
        <title>Manage Files</title>
        <meta name="description" content="View and manage your uploaded files" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex w-full flex-1 flex-col items-center justify-center px-20 text-center">
        <h1 className="text-6xl font-bold">Manage Files</h1>
        <p className="mt-3 text-2xl">View and manage your uploaded files below.</p>

        <ul className="mt-6 w-full max-w-4xl list-disc text-left sm:w-full">
          <li className="mt-2 text-xl">File 1</li>
          <li className="mt-2 text-xl">File 2</li>
          <li className="mt-2 text-xl">File 3</li>
        </ul>
      </main>

      <footer className="flex h-24 w-full items-center justify-center border-t">
        <Link href="/" className="flex items-center justify-center gap-2">
            Powered by{' '}
            <img src="/vercel.svg" alt="Vercel Logo" className="h-4" />
        </Link>
      </footer>
    </div>
  )
}
