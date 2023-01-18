/** Samples for MonitorOperation VMHostPayload. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/MainAccount_VMHosts_Payload.json
     */
    /**
     * Sample code: MainAccount_VMHosts_Payload.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void mainAccountVMHostsPayload(com.azure.resourcemanager.logz.LogzManager manager) {
        manager
            .monitorOperations()
            .vMHostPayloadWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
