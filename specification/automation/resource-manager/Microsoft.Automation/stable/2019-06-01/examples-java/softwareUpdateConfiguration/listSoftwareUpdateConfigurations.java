import com.azure.core.util.Context;

/** Samples for SoftwareUpdateConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/listSoftwareUpdateConfigurations.json
     */
    /**
     * Sample code: List software update configurations.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listSoftwareUpdateConfigurations(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.softwareUpdateConfigurations().listWithResponse("mygroup", "myaccount", null, null, Context.NONE);
    }
}
