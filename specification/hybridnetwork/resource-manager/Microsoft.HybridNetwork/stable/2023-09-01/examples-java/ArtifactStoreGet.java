
/**
 * Samples for ArtifactStores Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ArtifactStoreGet.
     * json
     */
    /**
     * Sample code: Get a artifact store resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getAArtifactStoreResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.artifactStores().getWithResponse("rg", "TestPublisher", "TestArtifactStoreName",
            com.azure.core.util.Context.NONE);
    }
}
