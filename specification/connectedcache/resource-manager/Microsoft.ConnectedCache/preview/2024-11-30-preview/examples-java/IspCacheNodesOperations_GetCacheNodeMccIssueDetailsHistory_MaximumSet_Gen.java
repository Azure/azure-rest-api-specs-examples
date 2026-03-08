
/**
 * Samples for IspCacheNodesOperations GetCacheNodeMccIssueDetailsHistory.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2024-11-30-preview/IspCacheNodesOperations_GetCacheNodeMccIssueDetailsHistory_MaximumSet_Gen.json
     */
    /**
     * Sample code: IspCacheNodesOperations_GetCacheNodeMccIssueDetailsHistory_MaximumSet.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void ispCacheNodesOperationsGetCacheNodeMccIssueDetailsHistoryMaximumSet(
        com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.ispCacheNodesOperations().getCacheNodeMccIssueDetailsHistoryWithResponse("rgConnectedCache",
            "MccRPTest1", "MCCCachenode1", com.azure.core.util.Context.NONE);
    }
}
