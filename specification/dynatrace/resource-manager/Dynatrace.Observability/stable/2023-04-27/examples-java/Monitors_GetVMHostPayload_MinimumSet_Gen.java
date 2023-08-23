/** Samples for Monitors GetVMHostPayload. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_GetVMHostPayload_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_GetVMHostPayload_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsGetVMHostPayloadMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .getVMHostPayloadWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
