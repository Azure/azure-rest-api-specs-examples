/** Samples for HostPools RetrieveRegistrationToken. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/HostPools_RetrieveRegistrationToken_Post.json
     */
    /**
     * Sample code: HostPools_RetrieveRegistrationToken_Post.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void hostPoolsRetrieveRegistrationTokenPost(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .hostPools()
            .retrieveRegistrationTokenWithResponse("resourceGroup1", "hostPool1", com.azure.core.util.Context.NONE);
    }
}
