import com.azure.core.util.Context;

/** Samples for HostPools ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/HostPool_ListByResourceGroup.json
     */
    /**
     * Sample code: HostPool_ListByResourceGroup.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void hostPoolListByResourceGroup(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.hostPools().listByResourceGroup("resourceGroup1", Context.NONE);
    }
}
