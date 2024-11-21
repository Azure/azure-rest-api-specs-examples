
/**
 * Samples for EnterpriseMccCustomers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/EnterpriseMccCustomers_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseMccCustomers_ListBySubscription.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void enterpriseMccCustomersListBySubscription(
        com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseMccCustomers().list(com.azure.core.util.Context.NONE);
    }
}
