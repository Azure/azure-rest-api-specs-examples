import com.azure.resourcemanager.desktopvirtualization.models.SessionHostPatch;

/** Samples for SessionHosts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/SessionHost_Update.json
     */
    /**
     * Sample code: SessionHost_Update.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void sessionHostUpdate(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .sessionHosts()
            .updateWithResponse(
                "resourceGroup1",
                "hostPool1",
                "sessionHost1.microsoft.com",
                true,
                new SessionHostPatch()
                    .withAllowNewSession(true)
                    .withAssignedUser("user1@microsoft.com")
                    .withFriendlyName("friendly"),
                com.azure.core.util.Context.NONE);
    }
}
