/** Samples for UserSessions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/UserSession_Delete.json
     */
    /**
     * Sample code: UserSession_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void userSessionDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .userSessions()
            .deleteWithResponse(
                "resourceGroup1",
                "hostPool1",
                "sessionHost1.microsoft.com",
                "1",
                true,
                com.azure.core.util.Context.NONE);
    }
}
