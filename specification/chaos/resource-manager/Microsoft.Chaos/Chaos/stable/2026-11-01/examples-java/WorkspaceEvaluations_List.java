
/**
 * Samples for WorkspaceEvaluations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceEvaluations_List.json
     */
    /**
     * Sample code: List retained workspace evaluations using canonical IDs and an opaque nextLink.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listRetainedWorkspaceEvaluationsUsingCanonicalIDsAndAnOpaqueNextLink(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceEvaluations().list("exampleRG", "exampleWorkspace", null, com.azure.core.util.Context.NONE);
    }
}
