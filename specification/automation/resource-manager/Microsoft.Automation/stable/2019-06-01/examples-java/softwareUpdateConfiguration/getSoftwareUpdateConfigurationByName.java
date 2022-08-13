import com.azure.core.util.Context;

/** Samples for SoftwareUpdateConfigurations GetByName. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/getSoftwareUpdateConfigurationByName.json
     */
    /**
     * Sample code: Get software update configuration by name.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getSoftwareUpdateConfigurationByName(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .softwareUpdateConfigurations()
            .getByNameWithResponse("mygroup", "myaccount", "mypatch", null, Context.NONE);
    }
}
