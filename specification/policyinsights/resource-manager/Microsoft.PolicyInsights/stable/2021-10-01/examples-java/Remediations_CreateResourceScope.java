import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.fluent.models.RemediationInner;

/** Samples for Remediations CreateOrUpdateAtResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_CreateResourceScope.json
     */
    /**
     * Sample code: Create remediation at individual resource scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void createRemediationAtIndividualResourceScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .remediations()
            .createOrUpdateAtResourceWithResponse(
                "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storAc1",
                "storageRemediation",
                new RemediationInner()
                    .withPolicyAssignmentId(
                        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/myResourceGroup/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
                Context.NONE);
    }
}
