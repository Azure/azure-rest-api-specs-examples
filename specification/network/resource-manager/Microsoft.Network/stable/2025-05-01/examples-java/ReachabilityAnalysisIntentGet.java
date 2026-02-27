
/**
 * Samples for ReachabilityAnalysisIntents Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ReachabilityAnalysisIntentGet
     * .json
     */
    /**
     * Sample code: ReachabilityAnalysisIntentGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reachabilityAnalysisIntentGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getReachabilityAnalysisIntents().getWithResponse("rg1",
            "testNetworkManager", "testWorkspace", "testAnalysisIntentName", com.azure.core.util.Context.NONE);
    }
}
