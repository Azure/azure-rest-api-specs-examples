
/**
 * Samples for Plans List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Plans_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Plans_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void
        plansListMinimumSetGen(com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.plans().list(null, "hilawwjz", com.azure.core.util.Context.NONE);
    }
}
