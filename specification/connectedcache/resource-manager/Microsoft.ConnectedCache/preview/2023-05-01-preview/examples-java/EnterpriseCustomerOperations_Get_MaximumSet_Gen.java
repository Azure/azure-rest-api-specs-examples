
/**
 * Samples for EnterpriseCustomerOperations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/EnterpriseCustomerOperations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseCustomerOperations_Get.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void
        enterpriseCustomerOperationsGet(com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseCustomerOperations().getByResourceGroupWithResponse("rgConnectedCache", "MCCTPTest2",
            com.azure.core.util.Context.NONE);
    }
}
