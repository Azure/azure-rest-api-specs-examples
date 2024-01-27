
import com.azure.core.util.Context;

/** Samples for ScheduledQueryRules GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/getScheduledQueryRule.json
     */
    /**
     * Sample code: Get a scheduled query rule for single resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAScheduledQueryRuleForSingleResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getScheduledQueryRules()
            .getByResourceGroupWithResponse("QueryResourceGroupName", "perf", Context.NONE);
    }
}
