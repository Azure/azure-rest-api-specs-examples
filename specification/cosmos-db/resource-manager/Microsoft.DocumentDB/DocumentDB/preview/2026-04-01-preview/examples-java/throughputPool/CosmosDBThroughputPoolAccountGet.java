
/**
 * Samples for ThroughputPoolAccount Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/throughputPool/CosmosDBThroughputPoolAccountGet.json
     */
    /**
     * Sample code: CosmosDB ThroughputPool Account Get.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBThroughputPoolAccountGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getThroughputPoolAccounts().getWithResponse("rgName", "tp1", "db1",
            com.azure.core.util.Context.NONE);
    }
}
