/** Samples for HostPools List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/HostPool_List.json
     */
    /**
     * Sample code: HostPool_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void hostPoolList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.hostPools().list(10, true, 0, com.azure.core.util.Context.NONE);
    }
}
