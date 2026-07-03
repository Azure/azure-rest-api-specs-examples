
/**
 * Samples for ReachabilityAnalysisIntents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ReachabilityAnalysisIntentGet.json
     */
    /**
     * Sample code: ReachabilityAnalysisIntentGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void reachabilityAnalysisIntentGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getReachabilityAnalysisIntents().getWithResponse("rg1", "testNetworkManager",
            "testWorkspace", "testAnalysisIntentName", com.azure.core.util.Context.NONE);
    }
}
