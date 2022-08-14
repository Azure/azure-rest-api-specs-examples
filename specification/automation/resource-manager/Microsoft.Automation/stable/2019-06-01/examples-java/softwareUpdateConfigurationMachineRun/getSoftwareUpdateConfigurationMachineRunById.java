import com.azure.core.util.Context;
import java.util.UUID;

/** Samples for SoftwareUpdateConfigurationMachineRuns GetById. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfigurationMachineRun/getSoftwareUpdateConfigurationMachineRunById.json
     */
    /**
     * Sample code: Get software update configuration machine run.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getSoftwareUpdateConfigurationMachineRun(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .softwareUpdateConfigurationMachineRuns()
            .getByIdWithResponse(
                "mygroup", "myaccount", UUID.fromString("ca440719-34a4-4234-a1a9-3f84faf7788f"), null, Context.NONE);
    }
}
