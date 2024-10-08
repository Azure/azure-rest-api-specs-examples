
/**
 * Samples for UserSessions ListByHostPool.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * UserSession_ListByHostPool.json
     */
    /**
     * Sample code: UserSession_ListByHostPool.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void userSessionListByHostPool(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.userSessions().listByHostPool("resourceGroup1", "hostPool1",
            "userPrincipalName eq 'user1@microsoft.com' and state eq 'active'", 10, true, 0,
            com.azure.core.util.Context.NONE);
    }
}
