import com.azure.core.util.Context;

/** Samples for Attestations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-01-01/examples/Attestations_ListSubscriptionScope_WithQuery.json
     */
    /**
     * Sample code: List attestations at subscription scope with query parameters.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void listAttestationsAtSubscriptionScopeWithQueryParameters(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .attestations()
            .list(
                1,
                "PolicyAssignmentId eq"
                    + " '/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5'"
                    + " AND PolicyDefinitionReferenceId eq '0b158b46-ff42-4799-8e39-08a5c23b4551'",
                Context.NONE);
    }
}
