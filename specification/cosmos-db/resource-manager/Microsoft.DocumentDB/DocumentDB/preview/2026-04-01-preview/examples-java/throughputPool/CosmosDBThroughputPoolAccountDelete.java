
/**
 * Samples for ThroughputPoolAccount Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/throughputPool/CosmosDBThroughputPoolAccountDelete.json
     */
    /**
     * Sample code: CosmosDB ThroughputPool Account Delete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBThroughputPoolAccountDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getThroughputPoolAccounts().delete("rgName", "tp1", "db1",
            com.azure.core.util.Context.NONE);
    }
}
