import com.azure.resourcemanager.newrelicobservability.models.MetricsStatusRequest;

/** Samples for Monitors GetMetricStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_GetMetricStatus_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_GetMetricStatus_MinimumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsGetMetricStatusMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager
            .monitors()
            .getMetricStatusWithResponse(
                "rgNewRelic",
                "fhcjxnxumkdlgpwanewtkdnyuz",
                new MetricsStatusRequest().withUserEmail("ruxvg@xqkmdhrnoo.hlmbpm"),
                com.azure.core.util.Context.NONE);
    }
}
