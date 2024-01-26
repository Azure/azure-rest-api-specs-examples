
import com.azure.core.util.Context;

/** Samples for MetricAlerts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getWebTestMetricAlert.json
     */
    /**
     * Sample code: Get a web test alert rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAWebTestAlertRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricAlerts()
            .getByResourceGroupWithResponse("rg-example", "webtest-name-example", Context.NONE);
    }
}
