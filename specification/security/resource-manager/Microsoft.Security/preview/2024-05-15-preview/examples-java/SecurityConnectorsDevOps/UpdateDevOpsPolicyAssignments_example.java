
import com.azure.resourcemanager.security.models.DescendantBehavior;
import com.azure.resourcemanager.security.models.DevOpsPolicyAssignment;
import com.azure.resourcemanager.security.models.DevOpsPolicyAssignmentProperties;
import com.azure.resourcemanager.security.models.DevOpsPolicyDescriptor;
import com.azure.resourcemanager.security.models.DevOpsPolicyType;

/**
 * Samples for DevOpsPolicyAssignments Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/UpdateDevOpsPolicyAssignments_example.json
     */
    /**
     * Sample code: Update_DevOpsPolicyAssignments.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void updateDevOpsPolicyAssignments(com.azure.resourcemanager.security.SecurityManager manager) {
        DevOpsPolicyAssignment resource
            = manager.devOpsPolicyAssignments().getWithResponse("myRg", "mySecurityConnectorName",
                "5ec87f43-62d8-437b-8f46-4c8d4032cf6d", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new DevOpsPolicyAssignmentProperties().withResourceId(
            "/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourcegroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/Contoso")
            .withDescendantBehavior(DescendantBehavior.OVERRIDE)
            .withPolicy(new DevOpsPolicyDescriptor().withPolicyName("myDevOpsPolicy")
                .withPolicyId("00000000-0000-0000-0000-000000000000").withPolicyVersion("1.0")
                .withPolicyType(DevOpsPolicyType.PIPELINE)))
            .apply();
    }
}
