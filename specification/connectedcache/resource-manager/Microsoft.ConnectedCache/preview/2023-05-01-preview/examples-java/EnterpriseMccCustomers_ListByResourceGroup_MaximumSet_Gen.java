
/**
 * Samples for EnterpriseMccCustomers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/EnterpriseMccCustomers_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseMccCustomers_ListByResourceGroup.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void enterpriseMccCustomersListByResourceGroup(
        com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseMccCustomers().listByResourceGroup("rgConnectedCache", com.azure.core.util.Context.NONE);
    }
}
