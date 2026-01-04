
/**
 * Samples for ReachabilityAnalysisRuns Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ReachabilityAnalysisRunGet.
     * json
     */
    /**
     * Sample code: ReachabilityAnalysisRunGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reachabilityAnalysisRunGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getReachabilityAnalysisRuns().getWithResponse("rg1",
            "testNetworkManager", "testWorkspace", "testAnalysisRunName", com.azure.core.util.Context.NONE);
    }
}
