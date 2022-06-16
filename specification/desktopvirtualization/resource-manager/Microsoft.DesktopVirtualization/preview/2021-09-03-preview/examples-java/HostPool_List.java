import com.azure.core.util.Context;

/** Samples for HostPools List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/HostPool_List.json
     */
    /**
     * Sample code: HostPool_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void hostPoolList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.hostPools().list(Context.NONE);
    }
}
