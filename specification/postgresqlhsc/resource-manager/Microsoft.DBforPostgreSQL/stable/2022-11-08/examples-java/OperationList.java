/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/OperationList.json
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
