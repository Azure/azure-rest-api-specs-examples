
/**
 * Samples for Monitors GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Monitors_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void
        monitorsGetMaximumSetGen(com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().getByResourceGroupWithResponse("rgNewRelic", "cdlymktqw", com.azure.core.util.Context.NONE);
    }
}
