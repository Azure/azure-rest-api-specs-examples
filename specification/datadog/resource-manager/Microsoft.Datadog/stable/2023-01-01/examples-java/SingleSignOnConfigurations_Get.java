/** Samples for SingleSignOnConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/SingleSignOnConfigurations_Get.json
     */
    /**
     * Sample code: SingleSignOnConfigurations_Get.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void singleSignOnConfigurationsGet(
        com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager
            .singleSignOnConfigurations()
            .getWithResponse("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE);
    }
}
