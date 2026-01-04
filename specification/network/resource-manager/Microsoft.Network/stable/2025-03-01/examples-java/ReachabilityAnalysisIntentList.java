
/**
 * Samples for ReachabilityAnalysisIntents List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ReachabilityAnalysisIntentList.json
     */
    /**
     * Sample code: ReachabilityAnalysisIntentList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reachabilityAnalysisIntentList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getReachabilityAnalysisIntents().list("rg1", "testNetworkManager",
            "testVerifierWorkspace1", null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
