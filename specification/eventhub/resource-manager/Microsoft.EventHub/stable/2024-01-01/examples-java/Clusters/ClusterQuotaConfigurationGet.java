
/**
 * Samples for Configuration Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/Clusters/
     * ClusterQuotaConfigurationGet.json
     */
    /**
     * Sample code: ClustersQuotasConfigurationGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void clustersQuotasConfigurationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getConfigurations().getWithResponse("myResourceGroup",
            "testCluster", com.azure.core.util.Context.NONE);
    }
}
