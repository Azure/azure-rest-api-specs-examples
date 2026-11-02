
/**
 * Samples for SvmPeers ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/SvmPeers_ListByParent.json
     */
    /**
     * Sample code: SvmPeers_ListByParent.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void svmPeersListByParent(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.svmPeers().listByParent("contosoNetAppRg", "contoso-external-cluster", "peerA",
            com.azure.core.util.Context.NONE);
    }
}
