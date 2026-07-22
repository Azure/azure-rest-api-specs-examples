
/**
 * Samples for RestorableTables List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBRestorableTableList.json
     */
    /**
     * Sample code: CosmosDBRestorableTableList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBRestorableTableList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getRestorableTables().list("WestUS", "98a570f2-63db-4117-91f0-366327b7b353", null, null,
            com.azure.core.util.Context.NONE);
    }
}
