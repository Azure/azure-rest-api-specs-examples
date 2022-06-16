import com.azure.core.util.Context;

/** Samples for MsixPackages Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/MsixPackage_Delete.json
     */
    /**
     * Sample code: MSIXPackage_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void mSIXPackageDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.msixPackages().deleteWithResponse("resourceGroup1", "hostpool1", "packagefullname", Context.NONE);
    }
}
