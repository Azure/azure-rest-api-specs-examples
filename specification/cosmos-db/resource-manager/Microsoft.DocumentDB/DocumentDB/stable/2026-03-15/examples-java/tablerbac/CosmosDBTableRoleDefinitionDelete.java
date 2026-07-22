
/**
 * Samples for TableResources DeleteTableRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/tablerbac/CosmosDBTableRoleDefinitionDelete.json
     */
    /**
     * Sample code: CosmosDBTableRoleDefinitionDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBTableRoleDefinitionDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getTableResources().deleteTableRoleDefinition("myResourceGroupName", "myAccountName",
            "myRoleDefinitionId", com.azure.core.util.Context.NONE);
    }
}
