
/**
 * Samples for WorkspaceDiscoveries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceDiscoveries_List.json
     */
    /**
     * Sample code: List retained workspace discoveries using canonical IDs and an opaque nextLink.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listRetainedWorkspaceDiscoveriesUsingCanonicalIDsAndAnOpaqueNextLink(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceDiscoveries().list("exampleRG", "exampleWorkspace", null, com.azure.core.util.Context.NONE);
    }
}
