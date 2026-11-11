
import com.azure.resourcemanager.chaos.models.ScenarioConfigurationProperties;

/**
 * Samples for ScenarioConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ScenarioConfigurations_CreateOrUpdate_CreateOnly.json
     */
    /**
     * Sample code: Create a scenario configuration only when absent, or return 412 when it already exists.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void createAScenarioConfigurationOnlyWhenAbsentOrReturn412WhenItAlreadyExists(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioConfigurations().define("exampleConfiguration")
            .withExistingScenario("exampleRG", "exampleWorkspace", "exampleScenario")
            .withProperties(new ScenarioConfigurationProperties().withScenarioId(
                "/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/exampleScenario"))
            .withIfNoneMatch("*").create();
    }
}
