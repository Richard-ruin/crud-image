<script>
    import { onMount } from 'svelte';
    import { getImages } from '../services/api';
    import ImageItem from './ImageItem.svelte';
    
    export let onRefresh = () => {};
    
    let images = [];
    let loading = true;
    let error = '';
    
    // Load images on component mount
    onMount(async () => {
      await loadImages();
    });
    
    // Fetch images from the API
    async function loadImages() {
      try {
        loading = true;
        error = '';
        images = await getImages();
      } catch (err) {
        error = 'Failed to load images. Please try again.';
        console.error(err);
      } finally {
        loading = false;
      }
    }
    
    // Handle image update
    function handleImageUpdated(event) {
      const updatedImage = event.detail;
      images = images.map(img => (img.id === updatedImage.id ? updatedImage : img));
    }
    
    // Handle image deletion
    function handleImageDeleted(event) {
      const deletedId = event.detail;
      images = images.filter(img => img.id !== deletedId);
    }
    
    // Handle refresh button click
    function handleRefresh() {
      loadImages();
      onRefresh();
    }
  </script>
  
  <div class="space-y-4">
    <div class="flex justify-between items-center">
      <h2 class="text-xl font-bold">Image Gallery</h2>
      <button
        on:click={handleRefresh}
        class="inline-flex items-center px-3 py-1 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Refresh
      </button>
    </div>
    
    {#if error}
      <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
        {error}
      </div>
    {/if}
    
    {#if loading}
      <div class="text-center py-10">
        <p class="text-gray-500">Loading images...</p>
      </div>
    {:else if images.length === 0}
      <div class="text-center py-10 bg-gray-50 rounded-lg">
        <p class="text-gray-500">No images uploaded yet. Upload your first image!</p>
      </div>
    {:else}
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        {#each images as image (image.id)}
          <ImageItem 
            {image} 
            on:imageUpdated={handleImageUpdated} 
            on:imageDeleted={handleImageDeleted} 
          />
        {/each}
      </div>
    {/if}
  </div>