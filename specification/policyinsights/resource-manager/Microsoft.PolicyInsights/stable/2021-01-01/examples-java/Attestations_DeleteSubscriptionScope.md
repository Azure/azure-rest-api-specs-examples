Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-policyinsights_1.0.0-beta.2/sdk/policyinsights/azure-resourcemanager-policyinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Attestations DeleteAtSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-01-01/examples/Attestations_DeleteSubscriptionScope.json
     */
    /**
     * Sample code: Delete attestation at subscription scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void deleteAttestationAtSubscriptionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.attestations().deleteAtSubscriptionWithResponse("790996e6-9871-4b1f-9cd9-ec42cd6ced1e", Context.NONE);
    }
}
```
