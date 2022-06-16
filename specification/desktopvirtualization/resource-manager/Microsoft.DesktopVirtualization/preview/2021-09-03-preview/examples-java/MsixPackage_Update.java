import com.azure.core.util.Context;
import com.azure.resourcemanager.desktopvirtualization.models.MsixPackage;

/** Samples for MsixPackages Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/MsixPackage_Update.json
     */
    /**
     * Sample code: MSIXPackage_Update.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void mSIXPackageUpdate(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        MsixPackage resource =
            manager
                .msixPackages()
                .getWithResponse("resourceGroup1", "hostpool1", "msixpackagefullname", Context.NONE)
                .getValue();
        resource.update().withIsActive(true).withIsRegularRegistration(false).withDisplayName("displayname").apply();
    }
}
