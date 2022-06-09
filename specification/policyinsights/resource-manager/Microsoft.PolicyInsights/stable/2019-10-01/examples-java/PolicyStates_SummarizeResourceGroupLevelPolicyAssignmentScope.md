```java
import com.azure.core.util.Context;

/** Samples for PolicyStates SummarizeForResourceGroupLevelPolicyAssignment. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeResourceGroupLevelPolicyAssignmentScope.json
     */
    /**
     * Sample code: Summarize at policy assignment scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void summarizeAtPolicyAssignmentScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .summarizeForResourceGroupLevelPolicyAssignmentWithResponse(
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                "myResourceGroup",
                "b7a1ca2596524e3ab19597f2",
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-policyinsights_1.0.0-beta.2/sdk/policyinsights/azure-resourcemanager-policyinsights/README.md) on how to add the SDK to your project and authenticate.
