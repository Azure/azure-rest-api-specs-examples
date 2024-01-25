
/**
 * Samples for ProvisionedClusterInstances ListAdminKubeconfig.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/
     * ProvisionedClusterInstanceListAdminKubeconfig.json
     */
    /**
     * Sample code: ListClusterAdminCredentials.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listClusterAdminCredentials(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.provisionedClusterInstances().listAdminKubeconfig(
            "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster",
            com.azure.core.util.Context.NONE);
    }
}
