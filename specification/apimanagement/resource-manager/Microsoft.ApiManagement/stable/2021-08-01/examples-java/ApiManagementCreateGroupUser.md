Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for GroupUser Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGroupUser.json
     */
    /**
     * Sample code: ApiManagementCreateGroupUser.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGroupUser(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .groupUsers()
            .createWithResponse("rg1", "apimService1", "tempgroup", "59307d350af58404d8a26300", Context.NONE);
    }
}
```
