/** Samples for HostPools GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/HostPool_Get.json
     */
    /**
     * Sample code: HostPool_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void hostPoolGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .hostPools()
            .getByResourceGroupWithResponse("resourceGroup1", "hostPool1", com.azure.core.util.Context.NONE);
    }
}
