Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PolicyAssignments GetById. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyAssignmentById.json
     */
    /**
     * Sample code: Retrieve a policy assignment by ID.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveAPolicyAssignmentByID(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyAssignments()
            .getByIdWithResponse(
                "providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyAssignments/LowCostStorage",
                Context.NONE);
    }
}
```
