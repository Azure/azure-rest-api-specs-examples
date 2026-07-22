
/**
 * Samples for CassandraResources GetCassandraRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/cassandrarbac/CosmosDBCassandraRoleDefinitionGet.json
     */
    /**
     * Sample code: CosmosDBCassandraRoleDefinitionGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraRoleDefinitionGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().getCassandraRoleDefinitionWithResponse("myResourceGroupName",
            "myAccountName", "myRoleDefinitionId", com.azure.core.util.Context.NONE);
    }
}
