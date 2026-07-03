
/**
 * Samples for ReachabilityAnalysisRuns List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ReachabilityAnalysisRunList.json
     */
    /**
     * Sample code: ReachabilityAnalysisRunList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void reachabilityAnalysisRunList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getReachabilityAnalysisRuns().list("rg1", "testNetworkManager",
            "testVerifierWorkspace1", null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
