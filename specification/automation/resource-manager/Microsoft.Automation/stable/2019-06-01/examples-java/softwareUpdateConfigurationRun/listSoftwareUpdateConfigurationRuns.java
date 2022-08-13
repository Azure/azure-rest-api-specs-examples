import com.azure.core.util.Context;

/** Samples for SoftwareUpdateConfigurationRuns List. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfigurationRun/listSoftwareUpdateConfigurationRuns.json
     */
    /**
     * Sample code: List software update configuration machine runs.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listSoftwareUpdateConfigurationMachineRuns(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .softwareUpdateConfigurationRuns()
            .listWithResponse("mygroup", "myaccount", null, null, null, null, Context.NONE);
    }
}
