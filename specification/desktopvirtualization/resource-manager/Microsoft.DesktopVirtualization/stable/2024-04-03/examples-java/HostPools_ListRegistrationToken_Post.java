
/**
 * Samples for HostPools ListRegistrationTokens.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * HostPools_ListRegistrationToken_Post.json
     */
    /**
     * Sample code: HostPools_ListRegistrationToken_Post.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void hostPoolsListRegistrationTokenPost(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.hostPools().listRegistrationTokensWithResponse("resourceGroup1", "hostPool1",
            com.azure.core.util.Context.NONE);
    }
}
