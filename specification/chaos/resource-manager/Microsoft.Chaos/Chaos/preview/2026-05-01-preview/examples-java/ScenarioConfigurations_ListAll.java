
/**
 * Samples for ScenarioConfigurations ListAll.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ScenarioConfigurations_ListAll.json
     */
    /**
     * Sample code: Get a list of scenario configurations.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAListOfScenarioConfigurations(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioConfigurations().listAll("exampleRG", "exampleWorkspace",
            "12345678-1234-1234-1234-123456789012", com.azure.core.util.Context.NONE);
    }
}
