
import com.azure.core.util.Context;

/** Samples for MetricAlerts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/
     * getDynamicMetricAlertSingleResource.json
     */
    /**
     * Sample code: Get a dynamic alert rule for single resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getADynamicAlertRuleForSingleResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricAlerts().getByResourceGroupWithResponse("gigtest",
            "chiricutin", Context.NONE);
    }
}
