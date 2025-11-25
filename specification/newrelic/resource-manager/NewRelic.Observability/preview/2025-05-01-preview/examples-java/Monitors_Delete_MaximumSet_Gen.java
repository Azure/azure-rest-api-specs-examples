
/**
 * Samples for Monitors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Monitors_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsDeleteMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().delete("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron", "ruxvg@xqkmdhrnoo.hlmbpm",
            com.azure.core.util.Context.NONE);
    }
}
