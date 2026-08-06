
/**
 * Samples for Workspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Workspaces_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_Get_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void workspacesGetMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("rgdiscovery", "7c14ca107f929876a0",
            com.azure.core.util.Context.NONE);
    }
}
