Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.resources.models.Identity;
import com.azure.resourcemanager.resources.models.PolicyAssignmentUpdate;
import com.azure.resourcemanager.resources.models.ResourceIdentityType;

/** Samples for PolicyAssignments UpdateById. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/updatePolicyAssignmentWithIdentityById.json
     */
    /**
     * Sample code: Update policy assignment with a managed identity by ID.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updatePolicyAssignmentWithAManagedIdentityByID(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyAssignments()
            .updateByIdWithResponse(
                "providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyAssignments/LowCostStorage",
                new PolicyAssignmentUpdate()
                    .withLocation("eastus")
                    .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED)),
                Context.NONE);
    }
}
```
