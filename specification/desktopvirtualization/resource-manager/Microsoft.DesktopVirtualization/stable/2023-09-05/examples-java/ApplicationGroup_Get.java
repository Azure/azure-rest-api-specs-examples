/** Samples for ApplicationGroups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ApplicationGroup_Get.json
     */
    /**
     * Sample code: ApplicationGroup_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void applicationGroupGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .applicationGroups()
            .getByResourceGroupWithResponse("resourceGroup1", "applicationGroup1", com.azure.core.util.Context.NONE);
    }
}
