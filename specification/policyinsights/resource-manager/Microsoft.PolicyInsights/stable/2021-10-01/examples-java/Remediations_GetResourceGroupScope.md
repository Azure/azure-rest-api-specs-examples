```java
import com.azure.core.util.Context;

/** Samples for Remediations GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_GetResourceGroupScope.json
     */
    /**
     * Sample code: Get remediation at resource group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void getRemediationAtResourceGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.remediations().getByResourceGroupWithResponse("myResourceGroup", "storageRemediation", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-policyinsights_1.0.0-beta.2/sdk/policyinsights/azure-resourcemanager-policyinsights/README.md) on how to add the SDK to your project and authenticate.
