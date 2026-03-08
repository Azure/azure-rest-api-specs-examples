
/**
 * Samples for EnterpriseMccCacheNodesOperations GetCacheNodeAutoUpdateHistory.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2024-11-30-preview/EnterpriseMccCacheNodesOperations_GetCacheNodeAutoUpdateHistory_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseMccCacheNodesOperations_GetCacheNodeAutoUpdateHistory_MaximumSet.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void enterpriseMccCacheNodesOperationsGetCacheNodeAutoUpdateHistoryMaximumSet(
        com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseMccCacheNodesOperations().getCacheNodeAutoUpdateHistoryWithResponse("rgConnectedCache",
            "MccRPTest1", "MCCCachenode1", com.azure.core.util.Context.NONE);
    }
}
