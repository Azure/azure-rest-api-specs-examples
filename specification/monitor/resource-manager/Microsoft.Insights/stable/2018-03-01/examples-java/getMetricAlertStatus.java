
import com.azure.core.util.Context;

/** Samples for MetricAlertsStatus List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getMetricAlertStatus.json
     */
    /**
     * Sample code: Get an alert rule status.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAnAlertRuleStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricAlertsStatus().listWithResponse("gigtest",
            "chiricutin", Context.NONE);
    }
}
