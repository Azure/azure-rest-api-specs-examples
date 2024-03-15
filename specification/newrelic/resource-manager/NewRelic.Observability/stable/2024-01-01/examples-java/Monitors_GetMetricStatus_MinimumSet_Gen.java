
import com.azure.resourcemanager.newrelicobservability.models.MetricsStatusRequest;
import java.util.Arrays;

/**
 * Samples for Monitors GetMetricStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/
     * Monitors_GetMetricStatus_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_GetMetricStatus_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsGetMetricStatusMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().getMetricStatusWithResponse("rgNewRelic", "fhcjxnxumkdlgpwanewtkdnyuz",
            new MetricsStatusRequest().withAzureResourceIds(Arrays.asList(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgNewRelic/providers/NewRelic.Observability/monitors/fhcjxnxumkdlgpwanewtkdnyuz"))
                .withUserEmail("ruxvg@xqkmdhrnoo.hlmbpm"),
            com.azure.core.util.Context.NONE);
    }
}
