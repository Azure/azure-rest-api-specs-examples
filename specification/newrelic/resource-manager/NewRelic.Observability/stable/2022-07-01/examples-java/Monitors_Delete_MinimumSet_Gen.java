/** Samples for Monitors Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_Delete_MinimumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsDeleteMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager
            .monitors()
            .delete(
                "rgopenapi",
                "ruxvg@xqkmdhrnoo.hlmbpm",
                "ipxmlcbonyxtolzejcjshkmlron",
                com.azure.core.util.Context.NONE);
    }
}
