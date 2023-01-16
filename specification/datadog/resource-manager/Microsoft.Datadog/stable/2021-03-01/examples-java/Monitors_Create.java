/** Samples for Monitors Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Monitors_Create.json
     */
    /**
     * Sample code: Monitors_Create.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsCreate(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager
            .monitors()
            .define("myMonitor")
            .withRegion((String) null)
            .withExistingResourceGroup("myResourceGroup")
            .create();
    }
}
