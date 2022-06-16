import com.azure.core.util.Context;

/** Samples for HostPools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/HostPool_Delete.json
     */
    /**
     * Sample code: HostPool_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void hostPoolDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.hostPools().deleteWithResponse("resourceGroup1", "hostPool1", true, Context.NONE);
    }
}
