
/**
 * Samples for UserSessions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * UserSession_List.json
     */
    /**
     * Sample code: UserSession_List.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void
        userSessionList(com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.userSessions().list("resourceGroup1", "hostPool1", "sessionHost1.microsoft.com", 10, true, 0,
            com.azure.core.util.Context.NONE);
    }
}
