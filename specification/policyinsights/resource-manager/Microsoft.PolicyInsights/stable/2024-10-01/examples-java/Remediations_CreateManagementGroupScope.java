
import com.azure.resourcemanager.policyinsights.fluent.models.RemediationInner;

/**
 * Samples for Remediations CreateOrUpdateAtManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * Remediations_CreateManagementGroupScope.json
     */
    /**
     * Sample code: Create remediation at management group scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void createRemediationAtManagementGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().createOrUpdateAtManagementGroupWithResponse("financeMg", "storageRemediation",
            new RemediationInner().withPolicyAssignmentId(
                "/providers/microsoft.management/managementGroups/financeMg/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
            com.azure.core.util.Context.NONE);
    }
}
