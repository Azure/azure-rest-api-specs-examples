import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByHostPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/PrivateLinkResources_ListByHostPool.json
     */
    /**
     * Sample code: PrivateLinkResources_ListByHostPool.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void privateLinkResourcesListByHostPool(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.privateLinkResources().listByHostPool("resourceGroup1", "hostPool1", Context.NONE);
    }
}
