```java
import com.azure.core.util.Context;
import java.time.OffsetDateTime;

/** Samples for PolicyStates SummarizeForManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeManagementGroupScope.json
     */
    /**
     * Sample code: Summarize at management group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void summarizeAtManagementGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .summarizeForManagementGroupWithResponse(
                "myManagementGroup",
                0,
                OffsetDateTime.parse("2019-10-05T18:00:00Z"),
                OffsetDateTime.parse("2019-10-06T18:00:00Z"),
                "PolicyDefinitionAction eq 'deny' or PolicyDefinitionAction eq 'audit'",
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-policyinsights_1.0.0-beta.2/sdk/policyinsights/azure-resourcemanager-policyinsights/README.md) on how to add the SDK to your project and authenticate.
