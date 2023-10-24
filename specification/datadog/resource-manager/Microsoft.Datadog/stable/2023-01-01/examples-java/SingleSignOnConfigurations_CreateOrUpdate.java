/** Samples for SingleSignOnConfigurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/SingleSignOnConfigurations_CreateOrUpdate.json
     */
    /**
     * Sample code: SingleSignOnConfigurations_CreateOrUpdate.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void singleSignOnConfigurationsCreateOrUpdate(
        com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager
            .singleSignOnConfigurations()
            .define("default")
            .withExistingMonitor("myResourceGroup", "myMonitor")
            .create();
    }
}
