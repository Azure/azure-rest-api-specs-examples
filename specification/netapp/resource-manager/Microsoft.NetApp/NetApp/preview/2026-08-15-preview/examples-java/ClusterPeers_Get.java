
/**
 * Samples for ClusterPeers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ClusterPeers_Get.json
     */
    /**
     * Sample code: ClusterPeers_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void clusterPeersGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.clusterPeers().getWithResponse("contosoNetAppRg", "contoso-external-cluster", "peerA",
            com.azure.core.util.Context.NONE);
    }
}
