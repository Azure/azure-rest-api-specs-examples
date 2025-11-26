
import com.azure.resourcemanager.newrelicobservability.models.MetricsRequest;

/**
 * Samples for Monitors GetMetricRules.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Monitors_GetMetricRules_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_GetMetricRules_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsGetMetricRulesMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().getMetricRulesWithResponse("rgNewRelic", "fhcjxnxumkdlgpwanewtkdnyuz",
            new MetricsRequest().withUserEmail("ruxvg@xqkmdhrnoo.hlmbpm"), com.azure.core.util.Context.NONE);
    }
}
