
/**
 * Samples for Organizations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Organizations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void organizationsListMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.organizations().list("ruxvg@xqkmdhrnoo.hlmbpm", "egh", com.azure.core.util.Context.NONE);
    }
}
