
/**
 * Samples for Monitors ListHosts.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/Hosts_List.json
     */
    /**
     * Sample code: Monitors_ListHosts.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsListHosts(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().listHosts("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
