/** Samples for MsixPackages Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/MsixPackage_Get.json
     */
    /**
     * Sample code: MSIXPackage_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void mSIXPackageGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .msixPackages()
            .getWithResponse("resourceGroup1", "hostpool1", "packagefullname", com.azure.core.util.Context.NONE);
    }
}
