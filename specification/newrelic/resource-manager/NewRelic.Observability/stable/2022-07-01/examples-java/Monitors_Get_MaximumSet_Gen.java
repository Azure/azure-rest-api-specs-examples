/** Samples for Monitors GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsGetMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().getByResourceGroupWithResponse("rgNewRelic", "cdlymktqw", com.azure.core.util.Context.NONE);
    }
}
