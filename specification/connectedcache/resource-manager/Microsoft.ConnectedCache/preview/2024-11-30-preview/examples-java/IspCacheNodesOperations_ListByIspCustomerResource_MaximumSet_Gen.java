
/**
 * Samples for IspCacheNodesOperations ListByIspCustomerResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-30-preview/IspCacheNodesOperations_ListByIspCustomerResource_MaximumSet_Gen.json
     */
    /**
     * Sample code: IspCacheNodesOperations_ListByIspCustomerResource.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void ispCacheNodesOperationsListByIspCustomerResource(
        com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.ispCacheNodesOperations().listByIspCustomerResource("rgConnectedCache", "MccRPTest1",
            com.azure.core.util.Context.NONE);
    }
}
