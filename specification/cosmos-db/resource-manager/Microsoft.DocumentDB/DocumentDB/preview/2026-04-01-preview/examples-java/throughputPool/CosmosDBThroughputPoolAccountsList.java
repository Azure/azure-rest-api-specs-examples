
/**
 * Samples for ThroughputPoolAccountsOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/throughputPool/CosmosDBThroughputPoolAccountsList.json
     */
    /**
     * Sample code: CosmosDB ThroughputPool Account List.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBThroughputPoolAccountList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getThroughputPoolAccountsOperations().list("rgName", "tp1",
            com.azure.core.util.Context.NONE);
    }
}
