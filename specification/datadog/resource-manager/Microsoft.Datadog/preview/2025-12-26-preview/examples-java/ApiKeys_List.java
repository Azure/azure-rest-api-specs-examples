
/**
 * Samples for Monitors ListApiKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/ApiKeys_List.json
     */
    /**
     * Sample code: Monitors_ListApiKeys.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsListApiKeys(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().listApiKeys("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
