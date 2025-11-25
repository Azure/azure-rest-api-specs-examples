
/**
 * Samples for Plans List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Plans_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Plans_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void
        plansListMaximumSetGen(com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.plans().list("pwuxgvrmkk", "hilawwjz", com.azure.core.util.Context.NONE);
    }
}
