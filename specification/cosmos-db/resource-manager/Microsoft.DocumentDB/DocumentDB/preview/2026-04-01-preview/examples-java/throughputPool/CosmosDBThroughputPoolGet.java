
/**
 * Samples for ThroughputPool GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/throughputPool/CosmosDBThroughputPoolGet.json
     */
    /**
     * Sample code: CosmosDB ThroughputPool Get.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBThroughputPoolGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getThroughputPools().getByResourceGroupWithResponse("rgName", "tp1",
            com.azure.core.util.Context.NONE);
    }
}
