Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-policyinsights_1.0.0-beta.2/sdk/policyinsights/azure-resourcemanager-policyinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Remediations CancelAtResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_CancelResourceGroupScope.json
     */
    /**
     * Sample code: Cancel a remediation at resource group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void cancelARemediationAtResourceGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().cancelAtResourceGroupWithResponse("myResourceGroup", "myRemediation", Context.NONE);
    }
}
```
