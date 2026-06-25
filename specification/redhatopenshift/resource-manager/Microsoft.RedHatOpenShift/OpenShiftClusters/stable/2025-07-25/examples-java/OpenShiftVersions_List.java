
/**
 * Samples for OpenShiftVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-25/OpenShiftVersions_List.json
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
