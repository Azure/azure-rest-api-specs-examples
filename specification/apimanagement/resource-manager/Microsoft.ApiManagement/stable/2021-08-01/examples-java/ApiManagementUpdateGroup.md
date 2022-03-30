Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.GroupContract;

/** Samples for Group Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateGroup.json
     */
    /**
     * Sample code: ApiManagementUpdateGroup.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateGroup(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        GroupContract resource =
            manager.groups().getWithResponse("rg1", "apimService1", "tempgroup", Context.NONE).getValue();
        resource.update().withDisplayName("temp group").withIfMatch("*").apply();
    }
}
```
