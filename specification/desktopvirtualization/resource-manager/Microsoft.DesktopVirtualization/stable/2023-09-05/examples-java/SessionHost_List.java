/** Samples for SessionHosts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/SessionHost_List.json
     */
    /**
     * Sample code: SessionHost_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void sessionHostList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.sessionHosts().list("resourceGroup1", "hostPool1", 10, true, 0, com.azure.core.util.Context.NONE);
    }
}
