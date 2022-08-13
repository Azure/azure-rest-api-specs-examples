import com.azure.core.util.Context;

/** Samples for SoftwareUpdateConfigurationRuns List. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfigurationRun/listFailedSoftwareUpdateConfigurationRuns.json
     */
    /**
     * Sample code: List software update configuration machine run with status equal to 'Failed'.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listSoftwareUpdateConfigurationMachineRunWithStatusEqualToFailed(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .softwareUpdateConfigurationRuns()
            .listWithResponse("mygroup", "myaccount", null, "properties/status eq 'Failed'", null, null, Context.NONE);
    }
}
