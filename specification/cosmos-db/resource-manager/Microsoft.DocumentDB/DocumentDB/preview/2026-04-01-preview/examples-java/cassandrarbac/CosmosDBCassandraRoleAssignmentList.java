
/**
 * Samples for CassandraResources ListCassandraRoleAssignments.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/cassandrarbac/CosmosDBCassandraRoleAssignmentList.json
     */
    /**
     * Sample code: CosmosDBCassandraRoleAssignmentList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraRoleAssignmentList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().listCassandraRoleAssignments("myResourceGroupName",
            "myAccountName", com.azure.core.util.Context.NONE);
    }
}
