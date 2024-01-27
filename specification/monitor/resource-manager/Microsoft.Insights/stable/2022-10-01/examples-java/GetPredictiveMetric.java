
import com.azure.core.util.Context;
import java.time.Duration;

/** Samples for PredictiveMetric Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/GetPredictiveMetric.json
     */
    /**
     * Sample code: Get Metric for data.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getMetricForData(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPredictiveMetrics().getWithResponse("myRG",
            "vmss1-Autoscale-775", "2021-10-14T22:00:00.000Z/2021-10-16T22:00:00.000Z", Duration.parse("PT1H"),
            "Microsoft.Compute/virtualMachineScaleSets", "PercentageCPU", "Total", Context.NONE);
    }
}
