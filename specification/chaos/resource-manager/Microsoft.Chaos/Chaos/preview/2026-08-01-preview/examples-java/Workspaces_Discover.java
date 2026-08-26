
/**
 * Samples for Workspaces Discover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/Workspaces_Discover.json
     */
    /**
     * Sample code: Trigger resource discovery for the workspace.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void triggerResourceDiscoveryForTheWorkspace(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().discover("exampleRG", "exampleWorkspace", com.azure.core.util.Context.NONE);
    }
}
