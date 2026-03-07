
/**
 * Samples for IspCacheNodesOperations GetCacheNodeAutoUpdateHistory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-30-preview/IspCacheNodesOperations_GetCacheNodeAutoUpdateHistory_MaximumSet_Gen.json
     */
    /**
     * Sample code: IspCacheNodesOperations_GetCacheNodeAutoUpdateHistory_MaximumSet.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void ispCacheNodesOperationsGetCacheNodeAutoUpdateHistoryMaximumSet(
        com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.ispCacheNodesOperations().getCacheNodeAutoUpdateHistoryWithResponse("rgConnectedCache", "MccRPTest1",
            "MCCCachenode1", com.azure.core.util.Context.NONE);
    }
}
