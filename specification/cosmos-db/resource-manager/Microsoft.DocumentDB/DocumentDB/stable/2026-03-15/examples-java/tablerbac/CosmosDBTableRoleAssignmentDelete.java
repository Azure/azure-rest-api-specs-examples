
/**
 * Samples for TableResources DeleteTableRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/tablerbac/CosmosDBTableRoleAssignmentDelete.json
     */
    /**
     * Sample code: CosmosDBTableRoleAssignmentDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBTableRoleAssignmentDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getTableResources().deleteTableRoleAssignment("myResourceGroupName", "myAccountName",
            "myRoleAssignmentId", com.azure.core.util.Context.NONE);
    }
}
