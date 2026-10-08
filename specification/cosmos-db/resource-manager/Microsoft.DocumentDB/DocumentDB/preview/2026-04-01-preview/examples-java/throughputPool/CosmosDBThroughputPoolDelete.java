
/**
 * Samples for ThroughputPool Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/throughputPool/CosmosDBThroughputPoolDelete.json
     */
    /**
     * Sample code: CosmosDB ThroughputPool Delete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBThroughputPoolDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getThroughputPools().delete("rgName", "tp1", com.azure.core.util.Context.NONE);
    }
}
