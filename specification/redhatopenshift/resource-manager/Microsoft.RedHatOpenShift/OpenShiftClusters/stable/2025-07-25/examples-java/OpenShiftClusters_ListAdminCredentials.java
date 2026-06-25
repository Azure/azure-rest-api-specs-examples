
/**
 * Samples for OpenShiftClusters ListAdminCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-25/OpenShiftClusters_ListAdminCredentials.json
     */
    /**
     * Sample code: Lists admin kubeconfig of an OpenShift cluster with the specified subscription, resource group and
     * resource name.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void listsAdminKubeconfigOfAnOpenShiftClusterWithTheSpecifiedSubscriptionResourceGroupAndResourceName(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.openShiftClusters().listAdminCredentialsWithResponse("resourceGroup", "resourceName",
            com.azure.core.util.Context.NONE);
    }
}
