
/**
 * Samples for OpenShiftClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-25/OpenShiftClusters_Get.json
     */
    /**
     * Sample code: Gets a OpenShift cluster with the specified subscription, resource group and resource name.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void getsAOpenShiftClusterWithTheSpecifiedSubscriptionResourceGroupAndResourceName(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.openShiftClusters().getByResourceGroupWithResponse("resourceGroup", "resourceName",
            com.azure.core.util.Context.NONE);
    }
}
