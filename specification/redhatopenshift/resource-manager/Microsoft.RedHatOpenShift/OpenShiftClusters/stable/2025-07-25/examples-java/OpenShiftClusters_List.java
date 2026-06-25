
/**
 * Samples for OpenShiftClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-25/OpenShiftClusters_List.json
     */
    /**
     * Sample code: Lists OpenShift clusters in the specified subscription.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void listsOpenShiftClustersInTheSpecifiedSubscription(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.openShiftClusters().list(com.azure.core.util.Context.NONE);
    }
}
