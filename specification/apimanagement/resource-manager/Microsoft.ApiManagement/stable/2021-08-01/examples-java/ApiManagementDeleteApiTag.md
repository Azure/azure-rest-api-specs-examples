```java
import com.azure.core.util.Context;

/** Samples for Tag DetachFromApi. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiTag.json
     */
    /**
     * Sample code: ApiManagementDeleteApiTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteApiTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tags()
            .detachFromApiWithResponse(
                "rg1", "apimService1", "59d5b28d1f7fab116c282650", "59d5b28e1f7fab116402044e", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
