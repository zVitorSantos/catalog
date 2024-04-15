<template>
    <div class="input-group">

        <input class="form-control border-end-0 border" type="search" v-model="searchQuery" placeholder="Search"
            id="example-search-input">
        <span class="input-group-append">
            <button class="btn btn-outline-secondary bg-white border-start-0 border-bottom-0 border ms-n5" type="button"
                @click="search">
                <i class="fa fa-search"></i>
            </button>
        </span>
    </div>
</template>

<script>
export default {
    data() {
        return {
            searchQuery: ''
        };
    },
    methods: {
        search() {
            // Perform search logic here
            console.log('Search query:', this.searchQuery);
        }
    }
};
</script>

<style scoped>
/* Add any custom styles here */
</style>