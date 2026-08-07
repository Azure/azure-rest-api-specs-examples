
/**
 * Samples for ChatModelDeployments ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/ChatModelDeployments_ListByWorkspace_MaximumSet_Gen.json
     */
    /**
     * Sample code: ChatModelDeployments_ListByWorkspace_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        chatModelDeploymentsListByWorkspaceMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.chatModelDeployments().listByWorkspace("rgdiscovery", "0f2d15df9509076ccf",
            com.azure.core.util.Context.NONE);
    }
}
