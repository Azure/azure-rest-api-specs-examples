import com.azure.core.util.Context;

/** Samples for Monitors ListUserRoles. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/MainAccount_listUserRoles.json
     */
    /**
     * Sample code: MainAccount_VMHosts_Update.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void mainAccountVMHostsUpdate(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.monitors().listUserRoles("myResourceGroup", "myMonitor", null, Context.NONE);
    }
}
