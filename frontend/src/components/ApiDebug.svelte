<script>
    import { onMount } from 'svelte';
    
    let apiStatus = 'Checking API connection...';
    let apiResult = '';
    let backendRunning = false;
    
    // Test direct fetch to backend
    async function testApiConnection() {
      try {
        const url = 'http://localhost:8080/api/images';
        apiStatus = `Testing connection to ${url}...`;
        
        const response = await fetch(url);
        apiStatus = `Response status: ${response.status}`;
        backendRunning = response.ok;
        
        if (response.ok) {
          const data = await response.json();
          apiResult = JSON.stringify(data, null, 2);
          return `Connected successfully. Status: ${response.status}`;
        } else {
          return `Connection failed. Status: ${response.status}`;
        }
      } catch (error) {
        apiStatus = `Connection error: ${error.message}`;
        return `Error: ${error.message}`;
      }
    }
    
    // Test simple upload
    async function testUpload() {
      try {
        const formData = new FormData();
        formData.append('title', 'Test Image');
        formData.append('description', 'This is a test image');
        
        // Create a simple test image (1x1 transparent pixel)
        const base64Data = 'iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg==';
        const byteString = atob(base64Data);
        const ab = new ArrayBuffer(byteString.length);
        const ia = new Uint8Array(ab);
        for (let i = 0; i < byteString.length; i++) {
          ia[i] = byteString.charCodeAt(i);
        }
        const blob = new Blob([ab], { type: 'image/png' });
        const file = new File([blob], 'test-image.png', { type: 'image/png' });
        
        formData.append('image', file);
        
        const response = await fetch('http://localhost:8080/api/images', {
          method: 'POST',
          body: formData
        });
        
        if (response.ok) {
          const result = await response.json();
          return `Upload successful! Image ID: ${result.id}`;
        } else {
          return `Upload failed. Status: ${response.status}`;
        }
      } catch (error) {
        return `Upload error: ${error.message}`;
      }
    }
    
    onMount(async () => {
      apiStatus = await testApiConnection();
    });
  </script>
  
  <div class="p-4 bg-yellow-100 rounded-lg mb-6 text-left">
    <h3 class="font-bold text-lg mb-2">API Connection Status</h3>
    <div class="mb-2">
      <span class="font-medium">Status:</span> 
      <span class={apiStatus.includes('Connected successfully') ? 'text-green-600' : 'text-red-600'}>
        {apiStatus}
      </span>
    </div>
    
    <div class="flex space-x-2">
      <button 
        on:click={testApiConnection} 
        class="bg-blue-500 hover:bg-blue-700 text-white font-medium py-1 px-4 rounded"
      >
        Test Connection
      </button>
      
      <button 
        on:click={async () => { apiStatus = await testUpload(); }}
        class="bg-green-500 hover:bg-green-700 text-white font-medium py-1 px-4 rounded"
        disabled={!backendRunning}
      >
        Test Upload
      </button>
    </div>
    
    {#if apiResult}
      <div class="mt-4">
        <h4 class="font-semibold">API Response:</h4>
        <pre class="bg-gray-100 p-2 rounded mt-2 overflow-auto max-h-64 text-sm">{apiResult}</pre>
      </div>
    {/if}
  </div>