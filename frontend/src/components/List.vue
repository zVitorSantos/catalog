<template>
    <div class="list-container">
        <div class="left-column">
            <div v-for="item in items" :key="item.id" class="item">
                <img :src="item.image" alt="Item Image" class="item-image" />
            </div>
        </div>
        <div class="right-column">
            <!-- Conteúdo da segunda coluna -->
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            items: Array.from({ length: 20 }, (_, i) => ({
                id: i + 1,
                title: `Item ${i + 1}`,
                description: `This is item ${i + 1}`,
                price: (i + 1) * 10
            }))
        };
    }
};
</script>

<style scoped>
.list-container {
    position: relative;
    top: 30%;
    left: 10%;
    right: 10%;
    bottom: 5%;
    overflow-y: scroll;
}

.left-column {
    width: 15%;
    float: left;
}

.right-column {
    width: 85%;
    float: left;
}

.item {
    margin-bottom: 1rem;
}

.item-image {
    width: 100%;
}
</style>