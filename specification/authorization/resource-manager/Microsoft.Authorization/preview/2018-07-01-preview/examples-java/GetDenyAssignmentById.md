Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DenyAssignments GetById. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-07-01-preview/examples/GetDenyAssignmentById.json
     */
    /**
     * Sample code: Get deny assignment by ID.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDenyAssignmentByID(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getDenyAssignments()
            .getByIdWithResponse(
                "subscriptions/subId/resourcegroups/rgname/providers/Microsoft.Authorization/denyAssignments/daId",
                Context.NONE);
    }
}
```
