
/**
 * Samples for EnterpriseCustomerOperations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/EnterpriseCustomerOperations_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseCustomerOperations_ListByResourceGroup.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void enterpriseCustomerOperationsListByResourceGroup(
        com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseCustomerOperations().listByResourceGroup("rgConnectedCache",
            com.azure.core.util.Context.NONE);
    }
}
