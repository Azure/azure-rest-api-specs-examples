
/**
 * Samples for AppAttachPackage ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * AppAttachPackage_ListByResourceGroup.json
     */
    /**
     * Sample code: AppAttachPackage_ListByResourceGroup.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void appAttachPackageListByResourceGroup(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.appAttachPackages().listByResourceGroup("resourceGroup1", "HostPoolName eq 'hostpool1'",
            com.azure.core.util.Context.NONE);
    }
}
