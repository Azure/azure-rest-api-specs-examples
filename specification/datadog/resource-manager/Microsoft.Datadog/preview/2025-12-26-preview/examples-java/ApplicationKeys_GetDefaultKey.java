
/**
 * Samples for Monitors GetDefaultApplicationKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/ApplicationKeys_GetDefaultKey.json
     */
    /**
     * Sample code: Monitors_GetDefaultApplicationKey.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void
        monitorsGetDefaultApplicationKey(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().getDefaultApplicationKeyWithResponse("myResourceGroup", "myMonitor",
            com.azure.core.util.Context.NONE);
    }
}
