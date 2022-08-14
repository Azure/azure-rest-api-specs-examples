import com.azure.core.util.Context;

/** Samples for SoftwareUpdateConfigurationMachineRuns List. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfigurationMachineRun/listSoftwareUpdateConfigurationMachineRunsByRun.json
     */
    /**
     * Sample code: List software update configuration machine runs for a specific software update configuration run.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listSoftwareUpdateConfigurationMachineRunsForASpecificSoftwareUpdateConfigurationRun(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .softwareUpdateConfigurationMachineRuns()
            .listWithResponse(
                "mygroup",
                "myaccount",
                null,
                "$filter=properties/correlationId eq 0b943e57-44d3-4f05-898c-6e92aa617e59",
                null,
                null,
                Context.NONE);
    }
}
