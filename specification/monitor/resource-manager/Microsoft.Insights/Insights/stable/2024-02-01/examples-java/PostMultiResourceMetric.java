
import java.time.Duration;

/**
 * Samples for Metrics ListAtSubscriptionScopePost.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01/PostMultiResourceMetric.json
     */
    /**
     * Sample code: Post request for subscription level metric data.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        postRequestForSubscriptionLevelMetricData(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetrics().listAtSubscriptionScopePostWithResponse("westus2",
            "2021-06-08T19:00:00Z/2021-06-12T01:00:00Z", Duration.parse("PT6H"), "Data Disk Max Burst IOPS", "count",
            10, "count desc", "LUN eq '0' and Microsoft.ResourceId eq '*'", null, "microsoft.compute/virtualmachines",
            true, false, "LUN", null, com.azure.core.util.Context.NONE);
    }
}
