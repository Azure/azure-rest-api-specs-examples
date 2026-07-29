
/**
 * Samples for Monitors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/Monitors_Delete.json
     */
    /**
     * Sample code: Monitors_Delete.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsDelete(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().delete("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
