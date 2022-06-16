import com.azure.core.util.Context;

/** Samples for MsixPackages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/MsixPackage_List.json
     */
    /**
     * Sample code: MSIXPackage_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void mSIXPackageList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.msixPackages().list("resourceGroup1", "hostpool1", Context.NONE);
    }
}
