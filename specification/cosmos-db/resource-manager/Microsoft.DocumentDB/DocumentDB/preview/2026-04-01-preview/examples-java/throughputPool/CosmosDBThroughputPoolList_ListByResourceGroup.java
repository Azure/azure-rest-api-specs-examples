
/**
 * Samples for ThroughputPoolsOperation ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/throughputPool/CosmosDBThroughputPoolList_ListByResourceGroup.json
     */
    /**
     * Sample code: CosmosDB ThroughputPool List by Resource Group.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBThroughputPoolListByResourceGroup(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getThroughputPoolsOperations().listByResourceGroup("rgName",
            com.azure.core.util.Context.NONE);
    }
}
