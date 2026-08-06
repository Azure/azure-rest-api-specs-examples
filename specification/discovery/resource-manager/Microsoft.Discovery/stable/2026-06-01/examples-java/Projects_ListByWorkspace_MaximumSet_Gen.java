
/**
 * Samples for Projects ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Projects_ListByWorkspace_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_ListByWorkspace_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void projectsListByWorkspaceMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.projects().listByWorkspace("rgdiscovery", "7712974a18ec06d5e6", com.azure.core.util.Context.NONE);
    }
}
