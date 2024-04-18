<template>
    <div class="list-container">
        <div v-for="item in items" :key="item.ref" class="item">
            <div class="left-column">
                <h2>{{ item.ref }}</h2>
                <p>R$ ****</p>
            </div>
            <div class="right-column">
                <p>{{ item.descricao }}</p>
                <p>{{ item.categoria_1 }}</p>
                <p>{{ item.categoria_2 }}</p>
                <p>{{ item.categoria_3 }}</p>
                <p>{{ item.complementos }}</p>
                <p>{{ item.material }}</p>
                <p>Peso: {{ item.peso }}</p>
                <p>Altura: {{ item.altura }}</p>
                <p>Largura: {{ item.largura }}</p>
                <p>Comprimento: {{ item.comprimento }}</p>
                <p>Matriz: {{ item.matriz ? 'Sim' : 'Não' }}</p>
                <p>Piloto: {{ item.piloto ? 'Sim' : 'Não' }}</p>
                <p>Desenho: {{ item.desenho ? 'Sim' : 'Não' }}</p>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            items: [],
            searchQuery: ''
        };
    },
    async created() {
        const response = await axios.get('http://catalog.c1omuu6kgyx2.us-east-2.rds.amazonaws.com/api/v1/products');
        this.items = response.data;
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
.list-container {
    display: flex;
    flex-direction: column;
    align-items: stretch;
    border: 1px solid #000;
    border-radius: 10px;
    overflow-y: scroll;
    max-height: 500px;
    margin-top: 10%;
    width: 85vw;
    margin: 0 auto;
    position: relative;
}

.item {
    display: flex;
    margin-bottom: 1rem;
    border-bottom: 1px solid #000;
}

.item:last-child {
    border-bottom: none;
}

.left-column {
    width: 15%;
    height: 30%;
}

.right-column {
    width: 85%;
    height: 30%;
}

.item-image {
    width: 100%;
    height: 100%;
}
</style>