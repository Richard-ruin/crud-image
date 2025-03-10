// Make sure this matches your Go backend URL exactly
export const API_URL = 'http://localhost:8080/api';

// Utility function to handle API responses
const handleResponse = async (response) => {
  // Log full response for debugging
  console.log(`API Response [${response.url}]: Status ${response.status}`);
  
  if (!response.ok) {
    // Try to get error message from response
    try {
      const errorData = await response.json();
      throw new Error(`API Error: ${errorData.error || response.statusText}`);
    } catch (jsonError) {
      throw new Error(`API Error: ${response.status} ${response.statusText}`);
    }
  }
  
  try {
    return await response.json();
  } catch (jsonError) {
    console.error('Error parsing JSON response:', jsonError);
    return null; // Return null for empty responses
  }
};

// Get all images
export const getImages = async () => {
  try {
    console.log('Fetching from:', `${API_URL}/images`);
    const response = await fetch(`${API_URL}/images`);
    return await handleResponse(response);
  } catch (error) {
    console.error('Error fetching images:', error);
    throw error;
  }
};

// Get single image by ID
export const getImage = async (id) => {
  try {
    const response = await fetch(`${API_URL}/images/${id}`);
    return await handleResponse(response);
  } catch (error) {
    console.error(`Error fetching image ${id}:`, error);
    throw error;
  }
};

// Upload a new image
export const uploadImage = async (formData) => {
  try {
    console.log('Uploading to:', `${API_URL}/images`);
    console.log('Form data entries:', [...formData.entries()].map(entry => {
      // Don't log the actual file content
      if (entry[0] === 'image') {
        return [entry[0], `File: ${entry[1].name} (${entry[1].size} bytes)`];
      }
      return entry;
    }));
    
    const response = await fetch(`${API_URL}/images`, {
      method: 'POST',
      body: formData
      // Note: Don't set Content-Type header when sending FormData
      // The browser will automatically set it with the correct boundary
    });
    
    return await handleResponse(response);
  } catch (error) {
    console.error('Error uploading image:', error);
    throw error;
  }
};

// Update image details
export const updateImage = async (id, formData) => {
  try {
    const response = await fetch(`${API_URL}/images/${id}`, {
      method: 'PUT',
      body: formData
    });
    
    return await handleResponse(response);
  } catch (error) {
    console.error(`Error updating image ${id}:`, error);
    throw error;
  }
};

// Delete an image
export const deleteImage = async (id) => {
  try {
    const response = await fetch(`${API_URL}/images/${id}`, {
      method: 'DELETE'
    });
    
    return await handleResponse(response);
  } catch (error) {
    console.error(`Error deleting image ${id}:`, error);
    throw error;
  }
};