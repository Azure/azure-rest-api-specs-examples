
/**
 * Samples for ThroughputPoolsOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/throughputPool/CosmosDBThroughputPoolList.json
     */
    /**
     * Sample code: CosmosDB ThroughputPool List.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBThroughputPoolList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getThroughputPoolsOperations().list(com.azure.core.util.Context.NONE);
    }
}
