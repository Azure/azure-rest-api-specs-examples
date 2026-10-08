
import com.azure.resourcemanager.cosmos.models.ThroughputPoolUpdate;

/**
 * Samples for ThroughputPool Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/throughputPool/CosmosDBThroughputPoolUpdate.json
     */
    /**
     * Sample code: CosmosDB ThroughputPool Update.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBThroughputPoolUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getThroughputPools().update("rg1", "tp1",
            new ThroughputPoolUpdate().withMaxThroughput(10000), com.azure.core.util.Context.NONE);
    }
}
