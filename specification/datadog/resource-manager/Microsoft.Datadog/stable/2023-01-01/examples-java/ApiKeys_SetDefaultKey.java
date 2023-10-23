/** Samples for Monitors SetDefaultKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/ApiKeys_SetDefaultKey.json
     */
    /**
     * Sample code: Monitors_SetDefaultKey.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsSetDefaultKey(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager
            .monitors()
            .setDefaultKeyWithResponse("myResourceGroup", "myMonitor", null, com.azure.core.util.Context.NONE);
    }
}
