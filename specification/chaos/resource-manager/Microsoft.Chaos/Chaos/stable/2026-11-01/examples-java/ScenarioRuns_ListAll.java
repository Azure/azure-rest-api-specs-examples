
/**
 * Samples for ScenarioRuns ListAll.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ScenarioRuns_ListAll.json
     */
    /**
     * Sample code: Get a list of scenario runs.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAListOfScenarioRuns(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioRuns().listAll("exampleRG", "exampleWorkspace", "12345678-1234-1234-1234-123456789012",
            com.azure.core.util.Context.NONE);
    }
}
