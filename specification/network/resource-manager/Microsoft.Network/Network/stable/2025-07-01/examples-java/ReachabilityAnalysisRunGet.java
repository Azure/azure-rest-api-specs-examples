
/**
 * Samples for ReachabilityAnalysisRuns Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ReachabilityAnalysisRunGet.json
     */
    /**
     * Sample code: ReachabilityAnalysisRunGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void reachabilityAnalysisRunGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getReachabilityAnalysisRuns().getWithResponse("rg1", "testNetworkManager",
            "testWorkspace", "testAnalysisRunName", com.azure.core.util.Context.NONE);
    }
}
