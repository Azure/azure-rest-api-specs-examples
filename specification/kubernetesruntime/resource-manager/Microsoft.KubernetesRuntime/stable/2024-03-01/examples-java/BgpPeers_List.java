
/**
 * Samples for BgpPeers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/BgpPeers_List.json
     */
    /**
     * Sample code: BgpPeers_List.
     * 
     * @param manager Entry point to ContainerOrchestratorRuntimeManager.
     */
    public static void bgpPeersList(
        com.azure.resourcemanager.containerorchestratorruntime.ContainerOrchestratorRuntimeManager manager) {
        manager.bgpPeers().list(
            "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1",
            com.azure.core.util.Context.NONE);
    }
}
