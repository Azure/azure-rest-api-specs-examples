
/**
 * Samples for SyncSets List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/SyncSets_List.json
     */
    /**
     * Sample code: Lists SyncSets that belong to that Azure Red Hat OpenShift Cluster.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void listsSyncSetsThatBelongToThatAzureRedHatOpenShiftCluster(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.syncSets().list("resourceGroup", "resourceName", com.azure.core.util.Context.NONE);
    }
}
