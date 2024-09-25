
/**
 * Samples for AppAttachPackage GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * AppAttachPackage_Get.json
     */
    /**
     * Sample code: AppAttachPackage_Get.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void
        appAttachPackageGet(com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.appAttachPackages().getByResourceGroupWithResponse("resourceGroup1", "packagefullname",
            com.azure.core.util.Context.NONE);
    }
}
