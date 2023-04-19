/** Samples for Monitors VmHostPayload. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_VmHostPayload_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_VmHostPayload_MaximumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsVmHostPayloadMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager
            .monitors()
            .vmHostPayloadWithResponse("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron", com.azure.core.util.Context.NONE);
    }
}
