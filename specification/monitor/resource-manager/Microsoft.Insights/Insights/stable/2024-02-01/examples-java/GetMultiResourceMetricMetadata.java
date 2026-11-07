
import com.azure.resourcemanager.monitor.models.MetricResultType;

/**
 * Samples for Metrics ListAtSubscriptionScope.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01/GetMultiResourceMetricMetadata.json
     */
    /**
     * Sample code: Get subscription level metric metadata.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getSubscriptionLevelMetricMetadata(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetrics().listAtSubscriptionScopeWithResponse("westus2",
            "2021-06-10T02:23:16.129Z/2021-06-12T02:23:16.129Z", null, "Data Disk Max Burst IOPS", null, null, null,
            "LUN eq '0'", MetricResultType.fromString("metadata"), "microsoft.compute/virtualmachines", null, null,
            null, com.azure.core.util.Context.NONE);
    }
}
