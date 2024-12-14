
/**
 * Samples for ProxyArtifact List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/PureProxyArtifact
     * /ArtifactList.json
     */
    /**
     * Sample code: List artifacts under an artifact store.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        listArtifactsUnderAnArtifactStore(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.proxyArtifacts().list("TestResourceGroup", "TestPublisher", "TestArtifactStoreName",
            com.azure.core.util.Context.NONE);
    }
}
