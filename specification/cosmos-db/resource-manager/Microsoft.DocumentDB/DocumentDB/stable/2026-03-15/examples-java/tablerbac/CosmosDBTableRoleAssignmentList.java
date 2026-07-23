
/**
 * Samples for TableResources ListTableRoleAssignments.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/tablerbac/CosmosDBTableRoleAssignmentList.json
     */
    /**
     * Sample code: CosmosDBTableRoleAssignmentList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBTableRoleAssignmentList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getTableResources().listTableRoleAssignments("myResourceGroupName", "myAccountName",
            com.azure.core.util.Context.NONE);
    }
}
