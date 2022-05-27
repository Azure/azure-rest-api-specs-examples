Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PolicyAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/deletePolicyAssignment.json
     */
    /**
     * Sample code: Delete a policy assignment.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAPolicyAssignment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyAssignments()
            .deleteWithResponse("subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "EnforceNaming", Context.NONE);
    }
}
```
