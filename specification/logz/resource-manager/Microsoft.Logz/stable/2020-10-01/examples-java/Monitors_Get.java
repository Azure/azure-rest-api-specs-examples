import com.azure.core.util.Context;

/** Samples for Monitors GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_Get.json
     */
    /**
     * Sample code: Monitors_Get.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void monitorsGet(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.monitors().getByResourceGroupWithResponse("myResourceGroup", "myMonitor", Context.NONE);
    }
}
