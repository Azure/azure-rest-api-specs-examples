/** Samples for Monitors Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_Create.json
     */
    /**
     * Sample code: Monitors_Create.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void monitorsCreate(com.azure.resourcemanager.logz.LogzManager manager) {
        manager
            .monitors()
            .define("myMonitor")
            .withRegion((String) null)
            .withExistingResourceGroup("myResourceGroup")
            .create();
    }
}
