
/**
 * Samples for Monitors List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/
     * Monitors_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().list(com.azure.core.util.Context.NONE);
    }
}
