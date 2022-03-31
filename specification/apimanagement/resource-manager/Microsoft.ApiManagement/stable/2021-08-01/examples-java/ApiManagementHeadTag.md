Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Tag GetEntityState. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadTag.json
     */
    /**
     * Sample code: ApiManagementHeadTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tags().getEntityStateWithResponse("rg1", "apimService1", "59306a29e4bbd510dc24e5f9", Context.NONE);
    }
}
```
