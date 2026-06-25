
/**
 * Samples for OpenShiftClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-25/OpenShiftClusters_ListByResourceGroup.json
     */
    /**
     * Sample code: Lists OpenShift clusters in the specified subscription and resource group.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void listsOpenShiftClustersInTheSpecifiedSubscriptionAndResourceGroup(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.openShiftClusters().listByResourceGroup("resourceGroup", com.azure.core.util.Context.NONE);
    }
}
