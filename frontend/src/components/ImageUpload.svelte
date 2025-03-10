<script>
    import { createEventDispatcher } from 'svelte';
    import { uploadImage } from '../services/api';
  
    const dispatch = createEventDispatcher();
    
    let title = '';
    let description = '';
    let fileInput;
    let isUploading = false;
    let previewUrl = '';
    let error = '';
  
    // Handle image preview
    function handleFileChange(event) {
      const file = event.target.files[0];
      if (!file) {
        previewUrl = '';
        return;
      }
  
      // Check if the file is an image
      if (!file.type.startsWith('image/')) {
        error = 'Please select an image file.';
        previewUrl = '';
        return;
      }
  
      error = '';
      previewUrl = URL.createObjectURL(file);
    }
  
    // Handle form submission
    async function handleSubmit() {
      error = '';
      if (!fileInput.files[0]) {
        error = 'Please select an image to upload.';
        return;
      }
  
      if (!title.trim()) {
        error = 'Please enter a title.';
        return;
      }
  
      try {
        isUploading = true;
        const formData = new FormData();
        formData.append('title', title);
        formData.append('description', description);
        formData.append('image', fileInput.files[0]);
  
        console.log('Uploading image...', {
          title,
          description,
          fileSize: fileInput.files[0].size,
          fileName: fileInput.files[0].name
        });
        
        const response = await uploadImage(formData);
        console.log('Upload response:', response);
        
        // Reset the form
        title = '';
        description = '';
        fileInput.value = '';
        previewUrl = '';
        
        // Notify parent component
        dispatch('imageUploaded', response);
      } catch (err) {
        error = 'Failed to upload image. Please try again.';
        console.error('Upload error:', err);
      } finally {
        isUploading = false;
      }
    }
  </script>
  
  <div class="bg-white p-6 rounded-lg shadow-md">
    <h2 class="text-xl font-bold mb-4">Upload New Image</h2>
    
    {#if error}
      <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
        {error}
      </div>
    {/if}
    
    <form on:submit|preventDefault={handleSubmit} class="space-y-4">
      <div>
        <label for="title" class="block text-sm font-medium text-gray-700">Title</label>
        <input
          id="title"
          type="text"
          bind:value={title}
          class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
          placeholder="Enter image title"
        />
      </div>
      
      <div>
        <label for="description" class="block text-sm font-medium text-gray-700">Description</label>
        <textarea
          id="description"
          bind:value={description}
          class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
          rows="3"
          placeholder="Enter image description"
        ></textarea>
      </div>
      
      <div>
        <label for="image" class="block text-sm font-medium text-gray-700">Image</label>
        <input
          id="image"
          type="file"
          accept="image/*"
          bind:this={fileInput}
          on:change={handleFileChange}
          class="mt-1 block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-sm file:font-semibold file:bg-indigo-50 file:text-indigo-700 hover:file:bg-indigo-100"
        />
      </div>
      
      {#if previewUrl}
        <div class="mt-2">
          <p class="text-sm text-gray-500">Preview:</p>
          <img src={previewUrl} alt="Preview" class="mt-2 h-64 w-auto object-cover rounded-md" />
        </div>
      {/if}
      
      <div>
        <button
          type="submit"
          disabled={isUploading}
          class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50"
        >
          {isUploading ? 'Uploading...' : 'Upload Image'}
        </button>
      </div>
    </form>
  </div>