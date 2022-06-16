import com.azure.core.util.Context;

/** Samples for Monitors ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_ListByResourceGroup.json
     */
    /**
     * Sample code: Monitors_ListByResourceGroup.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void monitorsListByResourceGroup(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.monitors().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
