
/**
 * Samples for OpenShiftVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/OpenShiftVersions_List.json
     */
    /**
     * Sample code: Lists all OpenShift versions available to install in the specified location.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void listsAllOpenShiftVersionsAvailableToInstallInTheSpecifiedLocation(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.openShiftVersions().list("location", com.azure.core.util.Context.NONE);
    }
}
