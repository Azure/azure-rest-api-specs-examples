import com.azure.core.util.Context;

/** Samples for ApplicationGroups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ApplicationGroup_Get.json
     */
    /**
     * Sample code: ApplicationGroup_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void applicationGroupGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.applicationGroups().getByResourceGroupWithResponse("resourceGroup1", "applicationGroup1", Context.NONE);
    }
}
