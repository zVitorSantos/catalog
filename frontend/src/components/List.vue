<template>
    <div class="list">
        <div class="list-container">
            <div v-for="item in items" :key="item.ref" class="item">
                <div class="image-column">
                    <img class="item-image" :src="item.imagem" alt="Imagem do item" />
                </div>
                <div class="info-column">
                    <div class="info-block">
                        <strong>Ref:</strong>
                        <p>{{ item.ref || 'N/A' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Descrição:</strong>
                        <p>{{ item.descricao || 'N/A' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Valor:</strong>
                        <p>R$ ****</p>
                    </div>
                </div>
                <div class="info-column">
                    <div class="info-block">
                        <strong>Categoria 1:</strong>
                        <p>{{ item.categoria_1 || 'N/A' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Categoria 2:</strong>
                        <p>{{ item.categoria_2 || 'N/A' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Categoria 3:</strong>
                        <p>{{ item.categoria_3 || 'N/A' }}</p>
                    </div>
                </div>
                <div class="info-column">
                    <div class="info-block">
                        <strong>Complementos:</strong>
                        <p>{{ item.complementos || 'N/A' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Material:</strong>
                        <p>{{ item.material || 'N/A' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Peso:</strong>
                        <p>{{ item.peso || 'N/A' }}</p>
                    </div>
                </div>
                <div class="info-column">
                    <div class="info-block">
                        <strong>Altura:</strong>
                        <p>{{ item.altura || 'N/A' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Largura:</strong>
                        <p>{{ item.largura || 'N/A' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Comprimento:</strong>
                        <p>{{ item.comprimento || 'N/A' }}</p>
                    </div>
                </div>
                <div class="info-column">
                    <div class="info-block">
                        <strong>Matriz:</strong>
                        <p>{{ item.matriz ? 'Sim' : 'Não' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Piloto:</strong>
                        <p>{{ item.piloto ? 'Sim' : 'Não' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Desenho:</strong>
                        <p>{{ item.desenho ? 'Sim' : 'Não' }}</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
//import axios from 'axios';

export default {
    data() {
        return {
            items: [
                {
                    ref: '1111',
                    imagem: 'https://picsum.photos/500/300',
                    descricao: 'Descrição padrão',
                    categoria_1: 'Enfeite',
                    categoria_2: 'Ornavi',
                    categoria_3: 'Coração',
                    complementos: 'Resina',
                    material: 'ABS',
                    peso: '5',
                    altura: '12',
                    largura: '11',
                    comprimento: '1',
                    matriz: true,
                    piloto: false,
                    desenho: true,
                },
                {
                    ref: '1111',
                    imagem: 'https://picsum.photos/600/300',
                    descricao: 'Descrição padrão',
                    categoria_1: 'Enfeite',
                    categoria_2: 'Ornavi',
                    categoria_3: 'Coração',
                    complementos: 'Resina',
                    material: 'ABS',
                    peso: '5',
                    altura: '12',
                    largura: '11',
                    comprimento: '1',
                    matriz: true,
                    piloto: false,
                    desenho: true,
                },
                {
                    ref: '1111',
                    imagem: 'https://picsum.photos/400/200',
                    descricao: 'Descrição padrão',
                    categoria_1: 'Enfeite',
                    categoria_2: 'Ornavi',
                    categoria_3: 'Coração',
                    complementos: 'Resina',
                    material: 'ABS',
                    peso: '5',
                    altura: '12',
                    largura: '11',
                    comprimento: '1',
                    matriz: true,
                    piloto: false,
                    desenho: true,
                },
                {
                    ref: '1111',
                    imagem: 'https://picsum.photos/200/400',
                    descricao: 'Descrição padrão',
                    categoria_1: 'Enfeite',
                    categoria_2: 'Ornavi',
                    categoria_3: 'Coração',
                    complementos: 'Resina',
                    material: 'ABS',
                    peso: '5',
                    altura: '12',
                    largura: '11',
                    comprimento: '1',
                    matriz: true,
                    piloto: false,
                    desenho: true,
                },
                {
                    ref: '1111',
                    imagem: 'https://picsum.photos/500/300',
                    descricao: 'Descrição padrão',
                    categoria_1: 'Enfeite',
                    categoria_2: 'Ornavi',
                    categoria_3: 'Coração',
                    complementos: 'Resina',
                    material: 'ABS',
                    peso: '5',
                    altura: '12',
                    largura: '11',
                    comprimento: '1',
                    matriz: true,
                    piloto: false,
                    desenho: true,
                },
                {
                    ref: '1111',
                    imagem: 'https://picsum.photos/600/400',
                    descricao: 'Descrição padrão',
                    categoria_1: 'Enfeite',
                    categoria_2: 'Ornavi',
                    categoria_3: 'Coração',
                    complementos: 'Resina',
                    material: 'ABS',
                    peso: '5',
                    altura: '12',
                    largura: '11',
                    comprimento: '1',
                    matriz: true,
                    piloto: false,
                    desenho: true,
                },
                {
                    ref: '1111',
                    imagem: 'https://picsum.photos/900/200',
                    descricao: 'Descrição padrão',
                    categoria_1: 'Enfeite',
                    categoria_2: 'Ornavi',
                    categoria_3: 'Coração',
                    complementos: 'Resina',
                    material: 'ABS',
                    peso: '5',
                    altura: '12',
                    largura: '11',
                    comprimento: '1',
                    matriz: true,
                    piloto: false,
                    desenho: true,
                },
                {
                    ref: '1111',
                    imagem: 'https://picsum.photos/200/900',
                    descricao: 'Descrição padrão',
                    categoria_1: 'Enfeite',
                    categoria_2: 'Ornavi',
                    categoria_3: 'Coração',
                    complementos: 'Resina',
                    material: 'ABS',
                    peso: '5',
                    altura: '12',
                    largura: '11',
                    comprimento: '1',
                    matriz: true,
                    piloto: false,
                    desenho: true,
                },
            ],
            searchQuery: ''
        };
    },
    // async created() {
    //     const response = await axios.get('http://catalog.c1omuu6kgyx2.us-east-2.rds.amazonaws.com/api/v1/products');
    //     this.items = response.data;
    // },
}
</script>

<style scoped>

.list {
    margin-top: 10%;

}

.list-container {
    display: flex;
    flex-direction: column;
    align-items: stretch;
    border: 1px solid #000;
    border-radius: 10px;
    overflow-y: scroll;
    max-height: 600px;
    margin-top: 17%;
    width: 85vw;
    margin: 0 auto;
    position: relative;
}

.item {
    display: flex;
    align-items: center;
    border-bottom: 1px solid #000;
    height: 200px;
}

.item:last-child {
    border-bottom: none;
}

.image-column {
    width: 40%;
    height: 100%;
    display: flex;
    align-items: center; 
    justify-content: center; 
}
.info-column {
    width: 20%;
    height: 100%;
    display: flex;
    flex-direction: column;
}

.info-block {
    margin-bottom: 1%;
}

.item-image {
    max-width: 90%; 
    max-height: 90%;
    margin: 10%; 
    object-fit: contain;
}

/* Estiliza a barra de rolagem */
.list-container::-webkit-scrollbar {
    width: 9px; 
}

/* Estiliza o fundo da barra de rolagem */
.list-container::-webkit-scrollbar-track {
    background: transparent;
    border-radius: 20px;
    border: 1px solid #000; 
    margin: 2px 0;
}

/* Estiliza o controle deslizante da barra de rolagem */
.list-container::-webkit-scrollbar-thumb {
    background-color: darkgrey; 
    border-radius: 20px; 
}

</style>