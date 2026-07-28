
/**
 * Samples for Monitors List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/Monitors_List.json
     */
    /**
     * Sample code: Monitors_List.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsList(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().list(com.azure.core.util.Context.NONE);
    }
}
