
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/
     * OperationList.json
     */
    /**
     * Sample code: List all available operations.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void listAllAvailableOperations(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
