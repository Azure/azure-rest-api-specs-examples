import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.fluent.models.AttestationInner;
import com.azure.resourcemanager.policyinsights.models.ComplianceState;

/** Samples for Attestations CreateOrUpdateAtSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-01-01/examples/Attestations_CreateSubscriptionScope.json
     */
    /**
     * Sample code: Create attestation at subscription scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void createAttestationAtSubscriptionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .attestations()
            .createOrUpdateAtSubscription(
                "790996e6-9871-4b1f-9cd9-ec42cd6ced1e",
                new AttestationInner()
                    .withPolicyAssignmentId(
                        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5")
                    .withComplianceState(ComplianceState.COMPLIANT),
                Context.NONE);
    }
}
