import com.azure.resourcemanager.newrelicobservability.models.MetricsStatusRequest;
import java.util.Arrays;

/** Samples for Monitors GetMetricStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_GetMetricStatus_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_GetMetricStatus_MaximumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsGetMetricStatusMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager
            .monitors()
            .getMetricStatusWithResponse(
                "rgNewRelic",
                "fhcjxnxumkdlgpwanewtkdnyuz",
                new MetricsStatusRequest()
                    .withAzureResourceIds(Arrays.asList("enfghpfw"))
                    .withUserEmail("ruxvg@xqkmdhrnoo.hlmbpm"),
                com.azure.core.util.Context.NONE);
    }
}
