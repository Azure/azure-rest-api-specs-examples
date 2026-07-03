
/**
 * Samples for ReachabilityAnalysisIntents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ReachabilityAnalysisIntentList.json
     */
    /**
     * Sample code: ReachabilityAnalysisIntentList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void reachabilityAnalysisIntentList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getReachabilityAnalysisIntents().list("rg1", "testNetworkManager",
            "testVerifierWorkspace1", null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
