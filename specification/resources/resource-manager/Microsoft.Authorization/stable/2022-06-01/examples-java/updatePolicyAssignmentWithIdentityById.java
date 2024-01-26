
import com.azure.resourcemanager.resources.models.Identity;
import com.azure.resourcemanager.resources.models.PolicyAssignmentUpdate;
import com.azure.resourcemanager.resources.models.ResourceIdentityType;

/** Samples for PolicyAssignments UpdateById. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/
     * updatePolicyAssignmentWithIdentityById.json
     */
    /**
     * Sample code: Update policy assignment with a managed identity by ID.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updatePolicyAssignmentWithAManagedIdentityByID(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyAssignments().updateByIdWithResponse(
            "providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyAssignments/LowCostStorage",
            new PolicyAssignmentUpdate().withLocation("eastus").withIdentity(
                new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED)),
            com.azure.core.util.Context.NONE);
    }
}
