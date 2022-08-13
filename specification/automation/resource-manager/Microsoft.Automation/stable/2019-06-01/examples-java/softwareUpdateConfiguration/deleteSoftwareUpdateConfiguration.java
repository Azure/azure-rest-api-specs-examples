import com.azure.core.util.Context;

/** Samples for SoftwareUpdateConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/deleteSoftwareUpdateConfiguration.json
     */
    /**
     * Sample code: Delete software update configuration.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteSoftwareUpdateConfiguration(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .softwareUpdateConfigurations()
            .deleteWithResponse("mygroup", "myaccount", "mypatch", null, Context.NONE);
    }
}
