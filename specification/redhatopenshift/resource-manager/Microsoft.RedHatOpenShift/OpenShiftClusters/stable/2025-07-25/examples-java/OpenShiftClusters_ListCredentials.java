
/**
 * Samples for OpenShiftClusters ListCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-25/OpenShiftClusters_ListCredentials.json
     */
    /**
     * Sample code: Lists credentials of an OpenShift cluster with the specified subscription, resource group and
     * resource name.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void listsCredentialsOfAnOpenShiftClusterWithTheSpecifiedSubscriptionResourceGroupAndResourceName(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.openShiftClusters().listCredentialsWithResponse("resourceGroup", "resourceName",
            com.azure.core.util.Context.NONE);
    }
}
