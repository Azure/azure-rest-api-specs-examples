
import com.azure.resourcemanager.newrelicobservability.models.SaaSData;

/**
 * Samples for Monitors LinkSaaS.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Monitors_LinkSaaS.json
     */
    /**
     * Sample code: Monitors_LinkSaaS.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void
        monitorsLinkSaaS(com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().linkSaaS("myResourceGroup", "myMonitor", new SaaSData().withSaaSResourceId(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgNewRelic/providers/Microsoft.SaaS/resources/abcd"),
            com.azure.core.util.Context.NONE);
    }
}
