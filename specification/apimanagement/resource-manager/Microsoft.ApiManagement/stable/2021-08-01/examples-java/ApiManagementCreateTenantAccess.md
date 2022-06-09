```java
import com.azure.resourcemanager.apimanagement.models.AccessIdName;

/** Samples for TenantAccess Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateTenantAccess.json
     */
    /**
     * Sample code: ApiManagementCreateTenantAccess.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateTenantAccess(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tenantAccess()
            .define(AccessIdName.ACCESS)
            .withExistingService("rg1", "apimService1")
            .withEnabled(true)
            .withIfMatch("*")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
