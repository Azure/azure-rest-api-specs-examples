
/**
 * Samples for EnterpriseMccCacheNodesOperations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/EnterpriseMccCacheNodesOperations_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseMccCacheNodesOperations Delete Operation - generated by [MaximumSet] rule.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void enterpriseMccCacheNodesOperationsDeleteOperationGeneratedByMaximumSetRule(
        com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseMccCacheNodesOperations().delete("rgConnectedCache",
            "hsincngmssuzeyispnufqwinpopadvhsbsemisguxgovwdjwviqexocelijvuyr", "vwtrhdoxvkrunpliwcao",
            com.azure.core.util.Context.NONE);
    }
}
