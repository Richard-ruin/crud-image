<script>
	import ImageUpload from './components/ImageUpload.svelte';
	import ImageList from './components/ImageList.svelte';
	import ApiDebug from './components/ApiDebug.svelte';
	
	let refreshList = false;
	
	// Trigger refresh when a new image is uploaded
	function handleImageUploaded() {
	  refreshList = !refreshList;
	}
  </script>
  
  <main class="min-h-screen bg-gray-100 py-8 px-4 sm:px-6 lg:px-8">
	<div class="max-w-6xl mx-auto">
	  <div class="text-center mb-8 p-4 bg-blue-100 rounded-lg">
		<h1 class="text-3xl font-bold text-blue-900">Image CRUD Application</h1>
		<p class="mt-2 text-blue-600">Upload, view, edit, and delete your images</p>
	  </div>
	  
	  <!-- API Debug Component -->
	  <ApiDebug />
	  
	  <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
		<div class="lg:col-span-1">
		  <ImageUpload on:imageUploaded={handleImageUploaded} />
		</div>
		
		<div class="lg:col-span-2">
		  <ImageList onRefresh={() => (refreshList = !refreshList)} />
		</div>
	  </div>
	</div>
  </main>