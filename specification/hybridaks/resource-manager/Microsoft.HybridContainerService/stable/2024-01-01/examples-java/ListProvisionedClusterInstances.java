
/**
 * Samples for ProvisionedClusterInstances List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/
     * ListProvisionedClusterInstances.json
     */
    /**
     * Sample code: ListProvisionedClusterInstances.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listProvisionedClusterInstances(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.provisionedClusterInstances().list(
            "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster",
            com.azure.core.util.Context.NONE);
    }
}
