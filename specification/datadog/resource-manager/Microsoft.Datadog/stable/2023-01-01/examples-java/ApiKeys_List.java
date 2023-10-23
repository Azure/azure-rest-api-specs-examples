/** Samples for Monitors ListApiKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/ApiKeys_List.json
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
