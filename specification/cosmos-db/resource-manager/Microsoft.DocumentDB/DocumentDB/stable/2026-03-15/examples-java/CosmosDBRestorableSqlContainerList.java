
/**
 * Samples for RestorableSqlContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBRestorableSqlContainerList.json
     */
    /**
     * Sample code: CosmosDBRestorableSqlContainerList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBRestorableSqlContainerList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getRestorableSqlContainers().list("WestUS", "98a570f2-63db-4117-91f0-366327b7b353",
            "3fu-hg==", null, null, com.azure.core.util.Context.NONE);
    }
}
