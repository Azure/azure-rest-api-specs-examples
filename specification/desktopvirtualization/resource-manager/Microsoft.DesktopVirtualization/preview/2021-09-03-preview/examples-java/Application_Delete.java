import com.azure.core.util.Context;

/** Samples for Applications Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Application_Delete.json
     */
    /**
     * Sample code: Application_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void applicationDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.applications().deleteWithResponse("resourceGroup1", "applicationGroup1", "application1", Context.NONE);
    }
}
