
/**
 * Samples for CassandraResources ListCassandraRoleDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/cassandrarbac/CosmosDBCassandraRoleDefinitionList.json
     */
    /**
     * Sample code: CosmosDBCassandraRoleDefinitionList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraRoleDefinitionList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().listCassandraRoleDefinitions("myResourceGroupName",
            "myAccountName", com.azure.core.util.Context.NONE);
    }
}
