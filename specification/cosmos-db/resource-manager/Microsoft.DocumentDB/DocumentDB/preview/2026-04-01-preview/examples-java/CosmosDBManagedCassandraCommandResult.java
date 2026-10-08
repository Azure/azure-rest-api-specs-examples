
/**
 * Samples for CassandraClusters GetCommandAsyncResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBManagedCassandraCommandResult.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraCommandResult.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBManagedCassandraCommandResult(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraClusters().getCommandAsyncResourceWithResponse("cassandra-prod-rg",
            "cassandra-prod", "318653d0-3da5-4814-b8f6-429f2af0b2a4", com.azure.core.util.Context.NONE);
    }
}
