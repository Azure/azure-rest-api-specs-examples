
/**
 * Samples for Monitors VmHostPayload.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Monitors_VmHostPayload_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_VmHostPayload_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsVmHostPayloadMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().vmHostPayloadWithResponse("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron",
            com.azure.core.util.Context.NONE);
    }
}
