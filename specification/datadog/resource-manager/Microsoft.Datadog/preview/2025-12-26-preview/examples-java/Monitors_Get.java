
/**
 * Samples for Monitors GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/Monitors_Get.json
     */
    /**
     * Sample code: Monitors_Get.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsGet(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().getByResourceGroupWithResponse("myResourceGroup", "myMonitor",
            com.azure.core.util.Context.NONE);
    }
}
