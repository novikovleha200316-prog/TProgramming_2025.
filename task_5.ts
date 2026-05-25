/**
 * Структура City (Город) - Вариант 17
 * 
 * Определяет минимум 3 переменных и 3 метода структуры.
 * Содержит метод-конструктор, возвращающий структуру.
 */

class City {
    name: string;       // Название города
    country: string;    // Страна
    population: number; // Население (тыс. чел.)
    area: number;       // Площадь (кв. км)
    constructor(name: string, country: string, population: number, area: number) {
        this.name = name;
        this.country = country;
        this.population = population;
        this.area = area;
    }
    printInfo(): void {
        console.log(`Город: ${this.name}`);
        console.log(`Страна: ${this.country}`);
        console.log(`Население: ${this.population} тыс. чел.`);
        console.log(`Площадь: ${this.area.toFixed(2)} кв. км`);
        console.log("-------------------------");
    }
    setPopulation(newPopulation: number): void {
        if (newPopulation >= 0) {
            const oldPopulation = this.population;
            this.population = newPopulation;
            console.log(`Население города ${this.name} обновлено: было ${oldPopulation} тыс. чел., стало ${this.population} тыс. чел.`);
        } else {
            console.log(`Ошибка: население не может быть отрицательным (${newPopulation})`);
        }
        console.log("-------------------------");
    }
    getPopulationDensity(): number {
        if (this.area <= 0) {
            console.log(`Ошибка: площадь города ${this.name} не может быть равна нулю или отрицательной`);
            return 0;
        }
        const density = (this.population * 1000) / this.area;
        return density;
    }
}
function NewCity(name: string, country: string, population: number, area: number): City {
    return new City(name, country, population, area);
}
function main(): void {
    console.log("========================================");
    console.log("СТРУКТУРА CITY (ГОРОД) - Вариант 17");
    console.log("========================================\n");
    const city = NewCity("Москва", "Россия", 12500, 2511.0);
    console.log("=== ИНФОРМАЦИЯ О ГОРОДЕ ===");
    city.printInfo();
    const density = city.getPopulationDensity();
    console.log(`Плотность населения: ${density.toFixed(2)} чел./кв. км`);
    console.log("-------------------------");
    console.log("=== ОБНОВЛЕНИЕ ДАННЫХ ===");
    city.setPopulation(13000);
    console.log("=== ОБНОВЛЕННАЯ ИНФОРМАЦИЯ ===");
    city.printInfo();
    const newDensity = city.getPopulationDensity();
    console.log(`Обновленная плотность населения: ${newDensity.toFixed(2)} чел./кв. км`);
    console.log("-------------------------");
    console.log("\n=== ДОПОЛНИТЕЛЬНАЯ ДЕМОНСТРАЦИЯ ===");
    const city2 = new City("Санкт-Петербург", "Россия", 5600, 1439.0);
    city2.printInfo();
    const density2 = city2.getPopulationDensity();
    console.log(`Плотность населения: ${density2.toFixed(2)} чел./кв. км`);
}
main();
export { City, NewCity };
