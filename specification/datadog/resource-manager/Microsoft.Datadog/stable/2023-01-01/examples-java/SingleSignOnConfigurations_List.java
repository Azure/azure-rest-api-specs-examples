/** Samples for SingleSignOnConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/SingleSignOnConfigurations_List.json
     */
    /**
     * Sample code: SingleSignOnConfigurations_List.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void singleSignOnConfigurationsList(
        com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.singleSignOnConfigurations().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
