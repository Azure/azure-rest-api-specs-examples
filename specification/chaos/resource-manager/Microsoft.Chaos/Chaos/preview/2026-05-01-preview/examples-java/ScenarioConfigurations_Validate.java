
/**
 * Samples for ScenarioConfigurations Validate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ScenarioConfigurations_Validate.json
     */
    /**
     * Sample code: Validate the given scenario configuration.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void validateTheGivenScenarioConfiguration(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioConfigurations().validate("exampleRG", "exampleWorkspace",
            "12345678-1234-1234-1234-123456789012", "config-5678-9012-3456-789012345678",
            com.azure.core.util.Context.NONE);
    }
}
