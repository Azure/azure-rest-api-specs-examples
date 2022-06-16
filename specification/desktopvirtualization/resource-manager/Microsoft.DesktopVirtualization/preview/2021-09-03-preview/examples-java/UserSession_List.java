import com.azure.core.util.Context;

/** Samples for UserSessions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/UserSession_List.json
     */
    /**
     * Sample code: UserSession_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void userSessionList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.userSessions().list("resourceGroup1", "hostPool1", "sessionHost1.microsoft.com", Context.NONE);
    }
}
