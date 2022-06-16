import com.azure.core.util.Context;

/** Samples for Desktops List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Desktop_List.json
     */
    /**
     * Sample code: Desktop_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void desktopList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.desktops().list("resourceGroup1", "applicationGroup1", Context.NONE);
    }
}
