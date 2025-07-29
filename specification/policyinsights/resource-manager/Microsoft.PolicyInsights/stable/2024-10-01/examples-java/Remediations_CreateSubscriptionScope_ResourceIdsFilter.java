
import com.azure.resourcemanager.policyinsights.fluent.models.RemediationInner;
import com.azure.resourcemanager.policyinsights.models.RemediationFilters;
import com.azure.resourcemanager.policyinsights.models.RemediationPropertiesFailureThreshold;
import com.azure.resourcemanager.policyinsights.models.ResourceDiscoveryMode;
import java.util.Arrays;

/**
 * Samples for Remediations CreateOrUpdateAtSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * Remediations_CreateSubscriptionScope_ResourceIdsFilter.json
     */
    /**
     * Sample code: Create remediation at subscription scope with resourceIds filter.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void createRemediationAtSubscriptionScopeWithResourceIdsFilter(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().createOrUpdateAtSubscriptionWithResponse("storageRemediation", new RemediationInner()
            .withPolicyAssignmentId(
                "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5")
            .withPolicyDefinitionReferenceId("8c8fa9e4")
            .withResourceDiscoveryMode(ResourceDiscoveryMode.EXISTING_NON_COMPLIANT)
            .withFilters(new RemediationFilters().withLocations(Arrays.asList("eastus", "westus"))
                .withResourceIds(Arrays.asList(
                    "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/res2627/providers/Microsoft.Storage/storageAccounts/sto1125",
                    "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/testcmk3/providers/Microsoft.Storage/storageAccounts/sto3699",
                    "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/res9407/providers/Microsoft.Storage/storageAccounts/sto8596",
                    "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/testcmk3/providers/Microsoft.Storage/storageAccounts/sto6637",
                    "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/res8186/providers/Microsoft.Storage/storageAccounts/sto834",
                    "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourceGroups/testcmk3/providers/Microsoft.Storage/storageAccounts/sto9174")))
            .withResourceCount(42).withParallelDeployments(6).withFailureThreshold(
                new RemediationPropertiesFailureThreshold().withPercentage(0.1F)),
            com.azure.core.util.Context.NONE);
    }
}
