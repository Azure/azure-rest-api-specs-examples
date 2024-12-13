
import com.azure.resourcemanager.datadog.models.DatadogSingleSignOnProperties;
import com.azure.resourcemanager.datadog.models.SingleSignOnStates;

/**
 * Samples for SingleSignOnConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/
     * SingleSignOnConfigurations_CreateOrUpdate.json
     */
    /**
     * Sample code: SingleSignOnConfigurations_CreateOrUpdate.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void
        singleSignOnConfigurationsCreateOrUpdate(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.singleSignOnConfigurations().define("default").withExistingMonitor("myResourceGroup", "myMonitor")
            .withProperties(new DatadogSingleSignOnProperties().withSingleSignOnState(SingleSignOnStates.ENABLE)
                .withEnterpriseAppId("00000000-0000-0000-0000-000000000000"))
            .create();
    }
}
