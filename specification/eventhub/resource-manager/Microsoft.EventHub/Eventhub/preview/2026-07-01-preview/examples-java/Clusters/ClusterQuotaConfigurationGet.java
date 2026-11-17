
/**
 * Samples for Configuration Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Clusters/ClusterQuotaConfigurationGet.json
     */
    /**
     * Sample code: ClustersQuotasConfigurationGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void clustersQuotasConfigurationGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getConfigurations().getWithResponse("myResourceGroup", "testCluster",
            com.azure.core.util.Context.NONE);
    }
}
