<template>
    <div class="list">
        <div class="list-container">
            <div v-for="item in items" :key="item.ref" class="card">
                <div class="image-column">
                    <img class="item-image" :src="item.imagem" alt="Imagem do item" />
                </div>
                <div class="info-column">
                    <div class="info-block">
                        <strong>Ref:</strong>
                        <p>{{ item.ref || 'N/A' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Medidas:</strong>
                        <p>{{ item.altura || 'N/A' }} x {{ item.largura || 'N/A' }} x {{ item.comprimento || 'N/A' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Peso:</strong>
                        <p>{{ item.peso || 'N/A' }}</p>
                    </div>
                    <div class="info-block">
                        <strong>Valor:</strong>
                        <p>R$ ****</p>
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
                    imagem: 'https://picsum.photos/610/390',
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
                {
                    ref: '1111',
                    imagem: 'https://picsum.photos/511/300',
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
                    imagem: 'https://picsum.photos/510/310',
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
                    imagem: 'https://picsum.photos/490/300',
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
                    imagem: 'https://picsum.photos/390/200',
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
                    imagem: 'https://picsum.photos/320/400',
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
                    imagem: 'https://picsum.photos/230/400',
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
    display: flex;
    justify-content: center;
    margin-top: 10%;
}

.list-container {
    display: grid;
    grid-template-columns: repeat(5, 1fr);
    gap: 1rem;
    margin: 0 auto;
    position: relative;
    justify-content: center;
    align-items: center;
}

.info-block {
    font-size: 0.8rem; 
}

.card {
    display: flex;
    flex-direction: column;
    border: 1px solid #000;
    border-radius: 10px;
    overflow: hidden;
    width: 200px; 
    height: 300px; 
}

.image-column {
    flex: 0.4;
    display: flex;
    align-items: center;
    justify-content: center;
    max-height: 40%;
}

.info-column {
    flex: 0.6;
    display: grid; 
    grid-template-columns: repeat(2, 1fr); 
    padding: 1rem;
}


.item-image {
    width: 100%; 
    height: 100%;
    object-fit: contain;
    padding: 5px;
}

</style>