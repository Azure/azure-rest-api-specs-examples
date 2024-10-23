
/**
 * Samples for BgpPeers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/BgpPeers_Get.json
     */
    /**
     * Sample code: BgpPeers_Get.
     * 
     * @param manager Entry point to ContainerOrchestratorRuntimeManager.
     */
    public static void bgpPeersGet(
        com.azure.resourcemanager.containerorchestratorruntime.ContainerOrchestratorRuntimeManager manager) {
        manager.bgpPeers().getWithResponse(
            "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1",
            "testpeer", com.azure.core.util.Context.NONE);
    }
}
