import com.azure.core.util.Context;

/** Samples for Remediations ListForManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_ListManagementGroupScope_WithQuery.json
     */
    /**
     * Sample code: List remediations at management group scope with query parameters.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listRemediationsAtManagementGroupScopeWithQueryParameters(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .remediations()
            .listForManagementGroup(
                "financeMg",
                1,
                "PolicyAssignmentId eq"
                    + " '/providers/microsoft.management/managementGroups/financeMg/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5'",
                Context.NONE);
    }
}
