
/**
 * Samples for Workspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Workspaces_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        workspacesListByResourceGroupMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.workspaces().listByResourceGroup("rgdiscovery", com.azure.core.util.Context.NONE);
    }
}
