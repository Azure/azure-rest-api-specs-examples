
import com.azure.resourcemanager.desktopvirtualization.models.AppAttachPackage;

/**
 * Samples for AppAttachPackage Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * AppAttachPackage_Update.json
     */
    /**
     * Sample code: AppAttachPackage_Update.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void
        appAttachPackageUpdate(com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        AppAttachPackage resource = manager.appAttachPackages()
            .getByResourceGroupWithResponse("resourceGroup1", "msixpackagefullname", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }
}
