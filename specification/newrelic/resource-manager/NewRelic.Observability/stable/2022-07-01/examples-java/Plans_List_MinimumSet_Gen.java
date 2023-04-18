/** Samples for Plans List. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Plans_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Plans_List_MinimumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void plansListMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.plans().list(null, null, com.azure.core.util.Context.NONE);
    }
}
