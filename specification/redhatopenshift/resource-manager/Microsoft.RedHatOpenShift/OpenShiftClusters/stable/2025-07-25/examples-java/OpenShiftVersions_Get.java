
/**
 * Samples for OpenShiftVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-25/OpenShiftVersions_Get.json
     */
    /**
     * Sample code: Gets an available OpenShift version to install in the specified location.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void getsAnAvailableOpenShiftVersionToInstallInTheSpecifiedLocation(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.openShiftVersions().getWithResponse("location", "4.14.40", com.azure.core.util.Context.NONE);
    }
}
