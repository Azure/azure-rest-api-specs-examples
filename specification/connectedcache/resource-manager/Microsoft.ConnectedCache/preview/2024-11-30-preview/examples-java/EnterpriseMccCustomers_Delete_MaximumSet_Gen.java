
/**
 * Samples for EnterpriseMccCustomers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-30-preview/EnterpriseMccCustomers_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseMccCustomers_Delete.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void
        enterpriseMccCustomersDelete(com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseMccCustomers().delete("rgConnectedCache", "zktb", com.azure.core.util.Context.NONE);
    }
}
