
/**
 * Samples for EnterpriseMccCustomers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/EnterpriseMccCustomers_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseMccCustomers_Get.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void
        enterpriseMccCustomersGet(com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.enterpriseMccCustomers().getByResourceGroupWithResponse("rgConnectedCache",
            "pvilvqkofbjbykupeewgvzlmjao", com.azure.core.util.Context.NONE);
    }
}
