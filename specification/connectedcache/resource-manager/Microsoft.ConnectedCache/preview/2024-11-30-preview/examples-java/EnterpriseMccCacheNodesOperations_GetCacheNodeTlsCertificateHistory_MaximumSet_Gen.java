
/**
 * Samples for EnterpriseMccCacheNodesOperations GetCacheNodeTlsCertificateHistory.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2024-11-30-preview/EnterpriseMccCacheNodesOperations_GetCacheNodeTlsCertificateHistory_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseMccCacheNodesOperations_GetCacheNodeTlsCertificateHistory_MaximumSet.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void enterpriseMccCacheNodesOperationsGetCacheNodeTlsCertificateHistoryMaximumSet(
        com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseMccCacheNodesOperations().getCacheNodeTlsCertificateHistoryWithResponse("rgConnectedCache",
            "MccRPTest1", "MCCCachenode1", com.azure.core.util.Context.NONE);
    }
}
