
/**
 * Samples for Cache CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateCache.json
     */
    /**
     * Sample code: ApiManagementCreateCache.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateCache(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.caches().define("c1").withExistingService("rg1", "apimService1")
            .withDescription("Redis cache instances in West India")
            .withConnectionString("apim.redis.cache.windows.net:6380,password=xc,ssl=True,abortConnect=False")
            .withUseFromLocation("default")
            .withResourceId(
                "https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redis/apimservice1")
            .create();
    }
}
