
/**
 * Samples for Monitors GetDefaultKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/ApiKeys_GetDefaultKey.json
     */
    /**
     * Sample code: Monitors_GetDefaultKey.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsGetDefaultKey(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().getDefaultKeyWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
