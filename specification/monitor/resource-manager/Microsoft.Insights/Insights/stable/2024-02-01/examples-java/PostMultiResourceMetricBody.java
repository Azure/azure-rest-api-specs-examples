
import com.azure.resourcemanager.monitor.models.SubscriptionScopeMetricsRequestBodyParameters;

/**
 * Samples for Metrics ListAtSubscriptionScopePost.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01/PostMultiResourceMetricBody.json
     */
    /**
     * Sample code: Post request for subscription level metric data using body params.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void postRequestForSubscriptionLevelMetricDataUsingBodyParams(
        com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetrics().listAtSubscriptionScopePostWithResponse("westus2", null, null, null, null,
            null, null, null, null, null, null, null, null,
            new SubscriptionScopeMetricsRequestBodyParameters()
                .withTimespan("2021-06-08T19:00:00Z/2021-06-12T01:00:00Z").withInterval("PT6H")
                .withMetricNames("Data Disk Max Burst IOPS").withAggregation("count")
                .withFilter("LUN eq '0' and Microsoft.ResourceId eq '*'").withTop(10).withOrderBy("count desc")
                .withRollUpBy("LUN").withMetricNamespace("microsoft.compute/virtualmachines")
                .withAutoAdjustTimegrain(true).withValidateDimensions(false),
            com.azure.core.util.Context.NONE);
    }
}
