
/**
 * Samples for CassandraResources DeleteCassandraRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/cassandrarbac/CosmosDBCassandraRoleDefinitionDelete.json
     */
    /**
     * Sample code: CosmosDBCassandraRoleDefinitionDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraRoleDefinitionDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().deleteCassandraRoleDefinition("myResourceGroupName",
            "myAccountName", "myRoleDefinitionId", com.azure.core.util.Context.NONE);
    }
}
