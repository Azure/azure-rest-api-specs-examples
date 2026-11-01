
/**
 * Samples for ExternalClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalClusters_Get.json
     */
    /**
     * Sample code: ExternalClusters_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void externalClustersGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalClusters().getByResourceGroupWithResponse("contosoNetAppRg", "contoso-external-cluster",
            com.azure.core.util.Context.NONE);
    }
}
