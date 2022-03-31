Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for UserGroup List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListUserGroups.json
     */
    /**
     * Sample code: ApiManagementListUserGroups.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListUserGroups(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.userGroups().list("rg1", "apimService1", "57681833a40f7eb6c49f6acf", null, null, null, Context.NONE);
    }
}
```
