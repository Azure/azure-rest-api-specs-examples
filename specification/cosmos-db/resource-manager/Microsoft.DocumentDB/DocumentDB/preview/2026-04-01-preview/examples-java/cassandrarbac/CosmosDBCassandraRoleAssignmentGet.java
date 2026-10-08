
/**
 * Samples for CassandraResources GetCassandraRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/cassandrarbac/CosmosDBCassandraRoleAssignmentGet.json
     */
    /**
     * Sample code: CosmosDBCassandraRoleAssignmentGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraRoleAssignmentGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().getCassandraRoleAssignmentWithResponse("myResourceGroupName",
            "myAccountName", "myRoleAssignmentId", com.azure.core.util.Context.NONE);
    }
}
