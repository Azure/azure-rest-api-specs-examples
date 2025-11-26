
/**
 * Samples for Monitors Resubscribe.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Monitors_Resubscribe.json
     */
    /**
     * Sample code: Monitors_Resubscribe.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void
        monitorsResubscribe(com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().resubscribe("myResourceGroup", "myMonitor", null, com.azure.core.util.Context.NONE);
    }
}
