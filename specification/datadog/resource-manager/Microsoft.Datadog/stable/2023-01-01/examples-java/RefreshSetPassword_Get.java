/** Samples for Monitors RefreshSetPasswordLink. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/RefreshSetPassword_Get.json
     */
    /**
     * Sample code: Monitors_RefreshSetPasswordLink.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsRefreshSetPasswordLink(
        com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager
            .monitors()
            .refreshSetPasswordLinkWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
