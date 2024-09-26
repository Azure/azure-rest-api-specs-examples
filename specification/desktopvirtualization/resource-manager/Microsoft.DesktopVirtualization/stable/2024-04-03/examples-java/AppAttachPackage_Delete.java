
/**
 * Samples for AppAttachPackage Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * AppAttachPackage_Delete.json
     */
    /**
     * Sample code: AppAttachPackage_Delete.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void
        appAttachPackageDelete(com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.appAttachPackages().deleteByResourceGroupWithResponse("resourceGroup1", "packagefullname",
            com.azure.core.util.Context.NONE);
    }
}
