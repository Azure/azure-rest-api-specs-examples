
/**
 * Samples for ReachabilityAnalysisIntents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ReachabilityAnalysisIntentDelete.json
     */
    /**
     * Sample code: ReachabilityAnalysisIntentDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void reachabilityAnalysisIntentDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getReachabilityAnalysisIntents().deleteWithResponse("rg1", "testNetworkManager",
            "testWorkspace", "testAnalysisIntent", com.azure.core.util.Context.NONE);
    }
}
