import com.azure.core.util.Context;

/** Samples for Applications Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Application_Get.json
     */
    /**
     * Sample code: Application_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void applicationGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.applications().getWithResponse("resourceGroup1", "applicationGroup1", "application1", Context.NONE);
    }
}
