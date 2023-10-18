/** Samples for SessionHosts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/SessionHost_Delete.json
     */
    /**
     * Sample code: SessionHost_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void sessionHostDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .sessionHosts()
            .deleteWithResponse(
                "resourceGroup1", "hostPool1", "sessionHost1.microsoft.com", true, com.azure.core.util.Context.NONE);
    }
}
