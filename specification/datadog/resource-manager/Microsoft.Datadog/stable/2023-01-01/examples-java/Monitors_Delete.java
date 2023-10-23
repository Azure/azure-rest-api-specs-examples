/** Samples for Monitors Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/Monitors_Delete.json
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
