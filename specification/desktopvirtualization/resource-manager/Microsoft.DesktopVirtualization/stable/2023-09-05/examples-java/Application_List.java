/** Samples for Applications List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/Application_List.json
     */
    /**
     * Sample code: Applications_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void applicationsList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .applications()
            .list("resourceGroup1", "applicationGroup1", 10, true, 0, com.azure.core.util.Context.NONE);
    }
}
