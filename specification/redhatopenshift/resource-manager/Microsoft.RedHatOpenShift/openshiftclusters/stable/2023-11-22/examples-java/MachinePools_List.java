
/**
 * Samples for MachinePools List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/MachinePools_List.json
     */
    /**
     * Sample code: Lists MachinePools that belong to that Azure Red Hat OpenShift Cluster.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void listsMachinePoolsThatBelongToThatAzureRedHatOpenShiftCluster(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.machinePools().list("resourceGroup", "resourceName", com.azure.core.util.Context.NONE);
    }
}
