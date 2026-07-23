
/**
 * Samples for TableResources ListTableRoleDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/tablerbac/CosmosDBTableRoleDefinitionList.json
     */
    /**
     * Sample code: CosmosDBTableRoleDefinitionList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBTableRoleDefinitionList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getTableResources().listTableRoleDefinitions("myResourceGroupName", "myAccountName",
            com.azure.core.util.Context.NONE);
    }
}
