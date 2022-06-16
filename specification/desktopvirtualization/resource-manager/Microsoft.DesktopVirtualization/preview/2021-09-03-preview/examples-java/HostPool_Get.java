import com.azure.core.util.Context;

/** Samples for HostPools GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/HostPool_Get.json
     */
    /**
     * Sample code: HostPool_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void hostPoolGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.hostPools().getByResourceGroupWithResponse("resourceGroup1", "hostPool1", Context.NONE);
    }
}
