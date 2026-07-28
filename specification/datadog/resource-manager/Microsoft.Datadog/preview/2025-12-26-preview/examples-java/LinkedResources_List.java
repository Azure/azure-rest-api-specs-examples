
/**
 * Samples for Monitors ListLinkedResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/LinkedResources_List.json
     */
    /**
     * Sample code: Monitors_ListLinkedResources.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsListLinkedResources(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().listLinkedResources("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
