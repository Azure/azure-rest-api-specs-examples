
/**
 * Samples for DatadogMonitorResources LatestLinkedSaaS.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/Monitors_LatestLinkedSaaS.json
     */
    /**
     * Sample code: Monitors_LatestLinkedSaaS.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsLatestLinkedSaaS(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.datadogMonitorResources().latestLinkedSaaSWithResponse("myResourceGroup", "myMonitor",
            com.azure.core.util.Context.NONE);
    }
}
