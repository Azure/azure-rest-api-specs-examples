import com.azure.core.util.Context;

/** Samples for MonitorOperation ListVmHostUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/MainAccount_VMHosts_Update.json
     */
    /**
     * Sample code: MainAccount_VMHosts_Update.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void mainAccountVMHostsUpdate(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.monitorOperations().listVmHostUpdate("myResourceGroup", "myMonitor", null, Context.NONE);
    }
}
