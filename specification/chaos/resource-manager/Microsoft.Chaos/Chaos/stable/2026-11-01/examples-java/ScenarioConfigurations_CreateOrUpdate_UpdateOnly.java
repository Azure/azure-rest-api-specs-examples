
import com.azure.resourcemanager.chaos.models.ScenarioConfiguration;
import com.azure.resourcemanager.chaos.models.ScenarioConfigurationProperties;

/**
 * Samples for ScenarioConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ScenarioConfigurations_CreateOrUpdate_UpdateOnly.json
     */
    /**
     * Sample code: Update a scenario configuration only when present, or return 412 when it does not exist.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void updateAScenarioConfigurationOnlyWhenPresentOrReturn412WhenItDoesNotExist(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        ScenarioConfiguration resource = manager.scenarioConfigurations().getWithResponse("exampleRG",
            "exampleWorkspace", "exampleScenario", "exampleConfiguration", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new ScenarioConfigurationProperties().withScenarioId(
            "/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/exampleScenario"))
            .withIfMatch("*").apply();
    }
}
