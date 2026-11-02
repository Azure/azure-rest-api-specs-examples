
/**
 * Samples for ExternalClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalClusters_Delete.json
     */
    /**
     * Sample code: ExternalClusters_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void externalClustersDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalClusters().delete("contosoNetAppRg", "contoso-external-cluster",
            com.azure.core.util.Context.NONE);
    }
}
