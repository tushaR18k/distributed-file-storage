
export interface File{
    name: string;
}

export async function fetchFiles(token: any): Promise<File[]>{
    const res = await fetch(`http://${process.env.NEXT_PUBLIC_AUTH_API_HOST}:${process.env.NEXT_PUBLIC_AUTH_API_PORT}/api/files`,{
        method: 'GET',
        headers:{
            'Authorization': `Bearer ${token}`
        },
    });
    if(res.status === 401){
        return [{name: "Unauthorized"}]
    }
    if(!res.ok){
        throw new Error('Failed to fetch uploaded files!');
    }
    const data = await  res.json();
    return data.map((item:any) => ({name: item}));
}