
/**
 * Samples for EnterpriseCustomerOperations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/EnterpriseCustomerOperations_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseCustomerOperations_ListBySubscription.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void enterpriseCustomerOperationsListBySubscription(
        com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseCustomerOperations().list(com.azure.core.util.Context.NONE);
    }
}
