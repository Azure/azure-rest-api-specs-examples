
/**
 * Samples for OpenShiftClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-25/OpenShiftClusters_Delete.json
     */
    /**
     * Sample code: Deletes a OpenShift cluster with the specified subscription, resource group and resource name.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void deletesAOpenShiftClusterWithTheSpecifiedSubscriptionResourceGroupAndResourceName(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.openShiftClusters().delete("resourceGroup", "resourceName", com.azure.core.util.Context.NONE);
    }
}
