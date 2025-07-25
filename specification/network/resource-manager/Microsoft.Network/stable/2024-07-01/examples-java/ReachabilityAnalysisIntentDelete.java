
/**
 * Samples for ReachabilityAnalysisIntents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * ReachabilityAnalysisIntentDelete.json
     */
    /**
     * Sample code: ReachabilityAnalysisIntentDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reachabilityAnalysisIntentDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getReachabilityAnalysisIntents().deleteWithResponse("rg1",
            "testNetworkManager", "testWorkspace", "testAnalysisIntent", com.azure.core.util.Context.NONE);
    }
}
