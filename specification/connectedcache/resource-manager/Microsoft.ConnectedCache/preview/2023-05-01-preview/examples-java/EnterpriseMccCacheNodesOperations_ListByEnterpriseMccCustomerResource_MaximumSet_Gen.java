
/**
 * Samples for EnterpriseMccCacheNodesOperations ListByEnterpriseMccCustomerResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2023-05-01-preview/EnterpriseMccCacheNodesOperations_ListByEnterpriseMccCustomerResource_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseMccCacheNodesOperations_ListByEnterpriseMccCustomerResource.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void enterpriseMccCacheNodesOperationsListByEnterpriseMccCustomerResource(
        com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseMccCacheNodesOperations().listByEnterpriseMccCustomerResource("rgConnectedCache", "syjjjzk",
            com.azure.core.util.Context.NONE);
    }
}
