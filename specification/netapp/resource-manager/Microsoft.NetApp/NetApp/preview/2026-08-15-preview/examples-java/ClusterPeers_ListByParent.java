
/**
 * Samples for ClusterPeers ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ClusterPeers_ListByParent.json
     */
    /**
     * Sample code: ClusterPeers_ListByParent.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void clusterPeersListByParent(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.clusterPeers().listByParent("contosoNetAppRg", "contoso-external-cluster",
            com.azure.core.util.Context.NONE);
    }
}
