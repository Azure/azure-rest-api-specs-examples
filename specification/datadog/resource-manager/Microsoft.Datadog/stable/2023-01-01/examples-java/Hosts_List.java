/** Samples for Monitors ListHosts. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/Hosts_List.json
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
