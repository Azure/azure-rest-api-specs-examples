/** Samples for Desktops Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/Desktop_Get.json
     */
    /**
     * Sample code: Desktop_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void desktopGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .desktops()
            .getWithResponse("resourceGroup1", "applicationGroup1", "SessionDesktop", com.azure.core.util.Context.NONE);
    }
}
