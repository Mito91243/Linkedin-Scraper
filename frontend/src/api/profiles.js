export async function fetchProfiles(companyName, categoryId) {
    const url = new URL('http://localhost/profiles/view');
    url.searchParams.append('company', companyName);
    url.searchParams.append('category', categoryId);
  
    const response = await fetch(url.toString());
    if (!response.ok) {
      throw new Error('Failed to fetch profiles');
    }
    const data = await response.json();
    //console.log('API Response:', data); // Log the API response
    return data;
  }