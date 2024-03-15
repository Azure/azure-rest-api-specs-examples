
/**
 * Samples for ConnectedPartnerResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/
     * ConnectedPartnerResources_List.json
     */
    /**
     * Sample code: ConnectedPartnerResources_List.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void connectedPartnerResourcesList(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.connectedPartnerResources().list("myResourceGroup", "myMonitor", null,
            com.azure.core.util.Context.NONE);
    }
}
