
/**
 * Samples for Scenarios ListAll.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Scenarios_ListAll.json
     */
    /**
     * Sample code: Get a list of scenarios.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAListOfScenarios(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarios().listAll("exampleRG", "exampleWorkspace", com.azure.core.util.Context.NONE);
    }
}
