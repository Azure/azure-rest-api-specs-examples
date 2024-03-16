
/**
 * Samples for Monitors ListLinkedResources.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/LinkedResources_List.
     * json
     */
    /**
     * Sample code: Monitors_ListLinkedResources.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsListLinkedResources(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().listLinkedResources("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
