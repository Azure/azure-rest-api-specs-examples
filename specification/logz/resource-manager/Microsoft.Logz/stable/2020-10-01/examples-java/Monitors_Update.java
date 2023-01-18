import com.azure.resourcemanager.logz.models.LogzMonitorResource;

/** Samples for Monitors Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_Update.json
     */
    /**
     * Sample code: Monitors_Update.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void monitorsUpdate(com.azure.resourcemanager.logz.LogzManager manager) {
        LogzMonitorResource resource =
            manager
                .monitors()
                .getByResourceGroupWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
