/** Samples for Monitors GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/Monitors_Get.json
     */
    /**
     * Sample code: Monitors_Get.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsGet(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager
            .monitors()
            .getByResourceGroupWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
