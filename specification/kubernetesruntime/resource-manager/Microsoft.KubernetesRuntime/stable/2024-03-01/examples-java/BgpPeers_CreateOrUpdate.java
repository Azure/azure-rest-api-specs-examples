
import com.azure.resourcemanager.containerorchestratorruntime.models.BgpPeerProperties;

/**
 * Samples for BgpPeers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/BgpPeers_CreateOrUpdate.json
     */
    /**
     * Sample code: BgpPeers_CreateOrUpdate.
     * 
     * @param manager Entry point to ContainerOrchestratorRuntimeManager.
     */
    public static void bgpPeersCreateOrUpdate(
        com.azure.resourcemanager.containerorchestratorruntime.ContainerOrchestratorRuntimeManager manager) {
        manager.bgpPeers().define("testpeer").withExistingResourceUri(
            "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1")
            .withProperties(new BgpPeerProperties().withMyAsn(64500).withPeerAsn(64501).withPeerAddress("10.0.0.1"))
            .create();
    }
}
