
/**
 * Samples for SvmPeers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/SvmPeers_Get.json
     */
    /**
     * Sample code: SvmPeers_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void svmPeersGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.svmPeers().getWithResponse("contosoNetAppRg", "contoso-external-cluster", "peerA", "contoso-svm-peer-a",
            com.azure.core.util.Context.NONE);
    }
}
