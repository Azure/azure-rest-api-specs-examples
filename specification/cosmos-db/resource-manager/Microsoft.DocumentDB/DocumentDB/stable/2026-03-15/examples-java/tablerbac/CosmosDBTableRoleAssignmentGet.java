
/**
 * Samples for TableResources GetTableRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/tablerbac/CosmosDBTableRoleAssignmentGet.json
     */
    /**
     * Sample code: CosmosDBTableRoleAssignmentGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBTableRoleAssignmentGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getTableResources().getTableRoleAssignmentWithResponse("myResourceGroupName",
            "myAccountName", "myRoleAssignmentId", com.azure.core.util.Context.NONE);
    }
}
