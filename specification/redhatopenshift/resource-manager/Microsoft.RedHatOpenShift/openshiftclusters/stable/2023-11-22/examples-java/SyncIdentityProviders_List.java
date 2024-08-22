
/**
 * Samples for SyncIdentityProviders List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/SyncIdentityProviders_List.json
     */
    /**
     * Sample code: Lists SyncIdentityProviders that belong to that Azure Red Hat OpenShift Cluster.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void listsSyncIdentityProvidersThatBelongToThatAzureRedHatOpenShiftCluster(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.syncIdentityProviders().list("resourceGroup", "resourceName", com.azure.core.util.Context.NONE);
    }
}
