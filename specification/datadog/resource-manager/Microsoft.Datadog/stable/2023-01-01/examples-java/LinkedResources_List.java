/** Samples for Monitors ListLinkedResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/LinkedResources_List.json
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
