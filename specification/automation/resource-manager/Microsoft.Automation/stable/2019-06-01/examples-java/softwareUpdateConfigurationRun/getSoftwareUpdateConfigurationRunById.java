
import java.util.UUID;

/**
 * Samples for SoftwareUpdateConfigurationRuns GetById.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/
     * softwareUpdateConfigurationRun/getSoftwareUpdateConfigurationRunById.json
     */
    /**
     * Sample code: Get software update configuration runs by Id.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void
        getSoftwareUpdateConfigurationRunsById(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.softwareUpdateConfigurationRuns().getByIdWithResponse("mygroup", "myaccount",
            UUID.fromString("2bd77cfa-2e9c-41b4-a45b-684a77cfeca9"), null, com.azure.core.util.Context.NONE);
    }
}
