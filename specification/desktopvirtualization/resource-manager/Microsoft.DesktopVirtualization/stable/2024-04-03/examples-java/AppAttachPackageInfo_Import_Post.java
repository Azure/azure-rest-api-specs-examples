
import com.azure.resourcemanager.desktopvirtualization.models.AppAttachPackageArchitectures;
import com.azure.resourcemanager.desktopvirtualization.models.ImportPackageInfoRequest;

/**
 * Samples for AppAttachPackageInfo ImportMethod.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * AppAttachPackageInfo_Import_Post.json
     */
    /**
     * Sample code: AppAttachPackageInfo_Import.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void appAttachPackageInfoImport(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.appAttachPackageInfoes().importMethod("resourceGroup1", "hostpool1", new ImportPackageInfoRequest()
            .withPath("imagepath").withPackageArchitecture(AppAttachPackageArchitectures.X64),
            com.azure.core.util.Context.NONE);
    }
}
