
/**
 * Samples for CacheNodesOperations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/CacheNodesOperations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: CacheNodesOperations_Get.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void cacheNodesOperationsGet(com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.cacheNodesOperations().getByResourceGroupWithResponse("rgConnectedCache",
            "nqoxkgorhuzbhjwcegymzqbeydzjupemekt", com.azure.core.util.Context.NONE);
    }
}
