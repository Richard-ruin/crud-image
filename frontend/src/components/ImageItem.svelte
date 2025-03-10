<script>
    import { createEventDispatcher } from 'svelte';
    import { deleteImage, updateImage } from '../services/api';
    
    export let image;
    
    const dispatch = createEventDispatcher();
    
    let isEditing = false;
    let isDeleting = false;
    let editTitle = '';
    let editDescription = '';
    let error = '';
    
    // Initialize edit form with current values
    function startEditing() {
      editTitle = image.title;
      editDescription = image.description;
      isEditing = true;
    }
    
    // Cancel editing
    function cancelEditing() {
      isEditing = false;
      error = '';
    }
    
    // Save edited image
    async function saveImage() {
      if (!editTitle.trim()) {
        error = 'Title cannot be empty';
        return;
      }
      
      try {
        const formData = new FormData();
        formData.append('title', editTitle);
        formData.append('description', editDescription);
        
        console.log('Updating image:', image.id, { editTitle, editDescription });
        const updatedImage = await updateImage(image.id, formData);
        console.log('Update response:', updatedImage);
        
        image = updatedImage;
        isEditing = false;
        error = '';
        
        dispatch('imageUpdated', updatedImage);
      } catch (err) {
        error = 'Failed to update image';
        console.error('Update error:', err);
      }
    }
    
    // Delete image
    async function handleDelete() {
      if (confirm('Are you sure you want to delete this image?')) {
        try {
          isDeleting = true;
          console.log('Deleting image:', image.id);
          const result = await deleteImage(image.id);
          console.log('Delete response:', result);
          
          dispatch('imageDeleted', image.id);
        } catch (err) {
          error = 'Failed to delete image';
          isDeleting = false;
          console.error('Delete error:', err);
        }
      }
    }
    
    // Format date for display
    function formatDate(dateString) {
      const date = new Date(dateString);
      return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
    }
  </script>
  
  <div class="bg-white rounded-lg shadow-md overflow-hidden">
    <img 
      src={'http://localhost:8080' + image.filePath} 
      alt={image.title}
      class="w-full h-48 object-cover"
    />
    
    {#if error}
      <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 text-sm">
        {error}
      </div>
    {/if}
    
    <div class="p-4">
      {#if isEditing}
        <div class="space-y-3">
          <div>
            <label for="edit-title" class="block text-sm font-medium text-gray-700">Title</label>
            <input
              id="edit-title"
              type="text"
              bind:value={editTitle}
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
            />
          </div>
          
          <div>
            <label for="edit-description" class="block text-sm font-medium text-gray-700">Description</label>
            <textarea
              id="edit-description"
              bind:value={editDescription}
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
              rows="3"
            ></textarea>
          </div>
          
          <div class="flex space-x-2">
            <button
              on:click={saveImage}
              class="inline-flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Save
            </button>
            <button
              on:click={cancelEditing}
              class="inline-flex justify-center py-2 px-4 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Cancel
            </button>
          </div>
        </div>
      {:else}
        <h3 class="text-lg font-bold text-gray-900">{image.title}</h3>
        
        {#if image.description}
          <p class="mt-1 text-gray-600">{image.description}</p>
        {/if}
        
        <div class="mt-2 text-xs text-gray-500">
          <p>Created: {formatDate(image.createdAt)}</p>
          <p>Updated: {formatDate(image.updatedAt)}</p>
        </div>
        
        <div class="mt-4 flex space-x-2">
          <button
            on:click={startEditing}
            class="px-3 py-1 text-sm text-indigo-600 hover:text-indigo-900 font-medium"
          >
            Edit
          </button>
          <button
            on:click={handleDelete}
            disabled={isDeleting}
            class="px-3 py-1 text-sm text-red-600 hover:text-red-900 font-medium disabled:opacity-50"
          >
            {isDeleting ? 'Deleting...' : 'Delete'}
          </button>
        </div>
      {/if}
    </div>
  </div>