Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-policyinsights_1.0.0-beta.2/sdk/policyinsights/azure-resourcemanager-policyinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.fluent.models.RemediationInner;
import com.azure.resourcemanager.policyinsights.models.RemediationFilters;
import com.azure.resourcemanager.policyinsights.models.RemediationPropertiesFailureThreshold;
import com.azure.resourcemanager.policyinsights.models.ResourceDiscoveryMode;
import java.util.Arrays;

/** Samples for Remediations CreateOrUpdateAtSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_CreateSubscriptionScope_AllProperties.json
     */
    /**
     * Sample code: Create remediation at subscription scope with all properties.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void createRemediationAtSubscriptionScopeWithAllProperties(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .remediations()
            .createOrUpdateAtSubscriptionWithResponse(
                "storageRemediation",
                new RemediationInner()
                    .withPolicyAssignmentId(
                        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5")
                    .withPolicyDefinitionReferenceId("8c8fa9e4")
                    .withResourceDiscoveryMode(ResourceDiscoveryMode.RE_EVALUATE_COMPLIANCE)
                    .withFilters(new RemediationFilters().withLocations(Arrays.asList("eastus", "westus")))
                    .withResourceCount(42)
                    .withParallelDeployments(6)
                    .withFailureThreshold(new RemediationPropertiesFailureThreshold().withPercentage(0.1f)),
                Context.NONE);
    }
}
```
