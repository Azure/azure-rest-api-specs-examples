/** Samples for Monitors GetDefaultKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/ApiKeys_GetDefaultKey.json
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
