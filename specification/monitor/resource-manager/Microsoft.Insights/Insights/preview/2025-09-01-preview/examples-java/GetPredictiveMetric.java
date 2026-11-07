
import java.time.Duration;

/**
 * Samples for PredictiveMetric Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/GetPredictiveMetric.json
     */
    /**
     * Sample code: Get Metric for data.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getMetricForData(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPredictiveMetrics().getWithResponse("myRG", "vmss1-Autoscale-775",
            "2021-10-14T22:00:00.000Z/2021-10-16T22:00:00.000Z", Duration.parse("PT1H"),
            "Microsoft.Compute/virtualMachineScaleSets", "PercentageCPU", "Total", com.azure.core.util.Context.NONE);
    }
}
