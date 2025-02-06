import { isNaN, startCase, toLower } from "lodash";
import { DatabaseType } from "../generated/graphql";

export function isNumeric(str: string) {
    return !isNaN(Number(str));
}

export function createStub(name: string) {
    return name.split(" ").map(word => word.toLowerCase()).join("-");
}

export function toTitleCase(str: string) {
    return startCase(toLower(str));
}

export function isDevelopment() {
    return process.env.NODE_ENV === "development"
}

export function isMarkdown(text: string): boolean {
    const markdownPatterns = [
        /^#{1,6}\s+/,
        /^\s*[-*+]\s+/,
        /^\d+\.\s+/,
        /\*\*[^*]+\*\*/,
        /_[^_]+_/,
        /!\[.*?\]\(.*?\)/,
        /\[.*?\]\(.*?\)/,
        /^>\s+/,
        /`{1,3}[^`]*`{1,3}/,
        /-{3,}/,
    ];
    
    return markdownPatterns.some(pattern => pattern.test(text));
}

export function isValidJSON(str: string): boolean {
    // this allows it to start showing intellisense when it starts with {
    // even if it is not valid - better UX?
    return str.startsWith("{");
}

export function isNoSQL(databaseType: string) {
    switch (databaseType) {
        case DatabaseType.MongoDb:
        case DatabaseType.Redis:
        case DatabaseType.ElasticSearch:
            return true;
    }
    return false;
}

export function getDatabaseStorageUnitLabel(databaseType: string | undefined, singular: boolean = false) {
    switch(databaseType) {
        case DatabaseType.ElasticSearch:
            return singular ? "Index" : "Indices";
        case DatabaseType.MongoDb:
            return singular ? "Collection" : "Collections";
        case DatabaseType.Redis:
            return singular ? "Key" : "Keys";
        case DatabaseType.MySql:
        case DatabaseType.Postgres:
        case DatabaseType.ClickHouse:
        case DatabaseType.Sqlite3:
            return singular ? "Table" : "Tables";
    }
    return singular ? "Storage Unit" : "Storage Units";
}

export function chooseRandomItems<T>(array: T[], n: number = 3): T[] {
    if (n > array.length) {
        throw new Error("n cannot be greater than the array length");
    }

    const shuffledArray = [...array];
    for (let i = shuffledArray.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [shuffledArray[i], shuffledArray[j]] = [shuffledArray[j], shuffledArray[i]];
    }

    return shuffledArray.slice(0, n);
}