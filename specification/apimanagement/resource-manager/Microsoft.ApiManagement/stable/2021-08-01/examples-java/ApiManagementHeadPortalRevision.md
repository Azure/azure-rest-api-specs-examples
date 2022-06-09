```java
import com.azure.core.util.Context;

/** Samples for PortalRevision GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadPortalRevision.json
     */
    /**
     * Sample code: ApiManagementHeadPortalRevision.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadPortalRevision(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.portalRevisions().getEntityTagWithResponse("rg1", "apimService1", "20201112101010", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
