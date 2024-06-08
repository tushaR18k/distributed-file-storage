
export interface File{
    name: string;
}

export async function fetchFiles(): Promise<File[]>{
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/files`);
    if(!res.ok){
        throw new Error('Failed to fetch uploaded files!');
    }
    const data = await  res.json();
    return data.map((item:any) => ({name: item}));
}