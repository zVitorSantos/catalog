<template>
    <div class="input-group rounded">
        <input type="search" class="form-control rounded" v-model="searchQuery" placeholder="Pesquisar" 
            aria-label="Pesquisar" aria-describedby="search-addon" />
        <span class="input-group-text border-0" id="search-addon">
            <i class="fas fa-search"></i>
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
body {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

.input-group {
    width: 50vw;
    min-width: 30rem;
    height: 3rem;
    position: absolute;
    top: 20%;
    left: 50%;
    transform: translate(-50%, -50%);
}

.form-control {
    border-width: 1px; 
    border-color: black;
}
</style>