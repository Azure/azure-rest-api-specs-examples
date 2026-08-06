
/**
 * Samples for ChatModelDeployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/ChatModelDeployments_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: ChatModelDeployments_Delete_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        chatModelDeploymentsDeleteMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.chatModelDeployments().delete("rgdiscovery", "b8412416e166a6c264", "aaf6134e93bb6af594",
            com.azure.core.util.Context.NONE);
    }
}
