
/**
 * Samples for ReachabilityAnalysisRuns List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ReachabilityAnalysisRunList.
     * json
     */
    /**
     * Sample code: ReachabilityAnalysisRunList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reachabilityAnalysisRunList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getReachabilityAnalysisRuns().list("rg1", "testNetworkManager",
            "testVerifierWorkspace1", null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
