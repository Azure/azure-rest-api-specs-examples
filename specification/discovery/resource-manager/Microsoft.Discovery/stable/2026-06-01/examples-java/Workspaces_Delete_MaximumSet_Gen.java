
/**
 * Samples for Workspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Workspaces_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_Delete_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void workspacesDeleteMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.workspaces().delete("rgdiscovery", "d70f215a9c483cbd5c", com.azure.core.util.Context.NONE);
    }
}
