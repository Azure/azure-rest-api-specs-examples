
/**
 * Samples for SyncSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/SyncSets_Get.json
     */
    /**
     * Sample code: Gets a SyncSet with the specified subscription, resource group and resource name.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void getsASyncSetWithTheSpecifiedSubscriptionResourceGroupAndResourceName(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.syncSets().getWithResponse("resourceGroup", "resourceName", "childResourceName",
            com.azure.core.util.Context.NONE);
    }
}
