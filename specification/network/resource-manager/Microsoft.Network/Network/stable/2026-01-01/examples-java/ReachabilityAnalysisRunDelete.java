
/**
 * Samples for ReachabilityAnalysisRuns Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ReachabilityAnalysisRunDelete.json
     */
    /**
     * Sample code: ReachabilityAnalysisRunDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void reachabilityAnalysisRunDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getReachabilityAnalysisRuns().delete("rg1", "testNetworkManager", "testWorkspace",
            "testAnalysisRun", com.azure.core.util.Context.NONE);
    }
}
