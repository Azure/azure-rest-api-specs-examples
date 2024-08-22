
/**
 * Samples for OpenShiftClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/OpenShiftClusters_List.json
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
