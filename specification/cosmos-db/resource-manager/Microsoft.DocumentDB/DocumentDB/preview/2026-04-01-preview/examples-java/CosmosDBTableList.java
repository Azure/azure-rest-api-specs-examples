
/**
 * Samples for TableResources ListTables.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBTableList.json
     */
    /**
     * Sample code: CosmosDBTableList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBTableList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getTableResources().listTables("rgName", "ddb1", com.azure.core.util.Context.NONE);
    }
}
