import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.fluent.models.RemediationInner;

/** Samples for Remediations CreateOrUpdateAtSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_CreateSubscriptionScope.json
     */
    /**
     * Sample code: Create remediation at subscription scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void createRemediationAtSubscriptionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .remediations()
            .createOrUpdateAtSubscriptionWithResponse(
                "storageRemediation",
                new RemediationInner()
                    .withPolicyAssignmentId(
                        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
                Context.NONE);
    }
}
