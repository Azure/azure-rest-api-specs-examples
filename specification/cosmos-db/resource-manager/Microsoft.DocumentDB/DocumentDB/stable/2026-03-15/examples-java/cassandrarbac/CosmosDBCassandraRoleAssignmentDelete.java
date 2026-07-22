
/**
 * Samples for CassandraResources DeleteCassandraRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/cassandrarbac/CosmosDBCassandraRoleAssignmentDelete.json
     */
    /**
     * Sample code: CosmosDBCassandraRoleAssignmentDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraRoleAssignmentDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().deleteCassandraRoleAssignment("myResourceGroupName",
            "myAccountName", "myRoleAssignmentId", com.azure.core.util.Context.NONE);
    }
}
