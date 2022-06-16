import com.azure.core.util.Context;

/** Samples for StartMenuItems List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/StartMenuItem_List.json
     */
    /**
     * Sample code: StartMenuItem_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void startMenuItemList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.startMenuItems().list("resourceGroup1", "applicationGroup1", Context.NONE);
    }
}
