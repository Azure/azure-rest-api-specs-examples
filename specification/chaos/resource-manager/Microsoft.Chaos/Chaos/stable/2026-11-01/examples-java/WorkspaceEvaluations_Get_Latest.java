
/**
 * Samples for WorkspaceEvaluations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceEvaluations_Get_Latest.json
     */
    /**
     * Sample code: Read latest as an alias while returning the canonical GA evaluation identity.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void readLatestAsAnAliasWhileReturningTheCanonicalGAEvaluationIdentity(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceEvaluations().getWithResponse("exampleRG", "exampleWorkspace", "latest",
            com.azure.core.util.Context.NONE);
    }
}
