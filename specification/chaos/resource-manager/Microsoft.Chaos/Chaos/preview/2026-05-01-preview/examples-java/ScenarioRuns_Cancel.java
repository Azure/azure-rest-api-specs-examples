
/**
 * Samples for ScenarioRuns Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ScenarioRuns_Cancel.json
     */
    /**
     * Sample code: Cancel a running scenario run.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void cancelARunningScenarioRun(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.scenarioRuns().cancel("exampleRG", "exampleWorkspace", "12345678-1234-1234-1234-123456789012",
            "abcd1234-5678-9012-3456-789012345678", com.azure.core.util.Context.NONE);
    }
}
