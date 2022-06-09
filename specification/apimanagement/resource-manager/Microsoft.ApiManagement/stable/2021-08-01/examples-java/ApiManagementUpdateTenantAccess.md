```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.AccessIdName;
import com.azure.resourcemanager.apimanagement.models.AccessInformationContract;

/** Samples for TenantAccess Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateTenantAccess.json
     */
    /**
     * Sample code: ApiManagementUpdateTenantAccess.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateTenantAccess(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        AccessInformationContract resource =
            manager.tenantAccess().getWithResponse("rg1", "apimService1", AccessIdName.ACCESS, Context.NONE).getValue();
        resource.update().withEnabled(true).withIfMatch("*").apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
