import com.azure.core.util.Context;

/** Samples for UserSessions ListByHostPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/UserSession_ListByHostPool.json
     */
    /**
     * Sample code: UserSession_ListByHostPool.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void userSessionListByHostPool(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .userSessions()
            .listByHostPool(
                "resourceGroup1",
                "hostPool1",
                "userPrincipalName eq 'user1@microsoft.com' and state eq 'active'",
                Context.NONE);
    }
}
