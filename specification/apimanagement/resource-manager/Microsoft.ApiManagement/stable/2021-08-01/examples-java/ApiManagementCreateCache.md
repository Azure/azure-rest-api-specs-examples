```java
/** Samples for Cache CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateCache.json
     */
    /**
     * Sample code: ApiManagementCreateCache.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateCache(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .caches()
            .define("c1")
            .withExistingService("rg1", "apimService1")
            .withDescription("Redis cache instances in West India")
            .withConnectionString("apim.redis.cache.windows.net:6380,password=xc,ssl=True,abortConnect=False")
            .withUseFromLocation("default")
            .withResourceId(
                "https://management.azure.com/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redis/apimservice1")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
