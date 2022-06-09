```java
import com.azure.core.util.Context;

/** Samples for Cache ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListCaches.json
     */
    /**
     * Sample code: ApiManagementListCaches.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListCaches(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.caches().listByService("rg1", "apimService1", null, null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
