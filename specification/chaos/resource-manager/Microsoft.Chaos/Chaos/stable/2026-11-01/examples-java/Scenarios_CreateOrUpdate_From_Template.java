
import com.azure.resourcemanager.chaos.models.ScenarioProperties;

/**
 * Samples for Scenarios CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Scenarios_CreateOrUpdate_From_Template.json
     */
    /**
     * Sample code: Create a scenario from a built-in scenario template version.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        createAScenarioFromABuiltInScenarioTemplateVersion(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarios().define("computeZoneDownScenario").withExistingWorkspace("exampleRG", "exampleWorkspace")
            .withProperties(new ScenarioProperties()
                .withCreatedFrom("/providers/Microsoft.Chaos/scenarioTemplates/ComputeZoneDown/versions/1.2")
                .withDescription(
                    "Simulates a datacenter/availability zone failure by shutting down VMSS instances and standalone VMs in the target zone. Zone failures account for ~25% of all cloud incidents — this scenario validates your application's cross-zone resilience."))
            .create();
    }
}
