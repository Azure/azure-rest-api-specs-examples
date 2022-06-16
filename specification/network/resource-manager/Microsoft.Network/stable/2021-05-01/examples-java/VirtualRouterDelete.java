import com.azure.core.util.Context;

/** Samples for VirtualRouters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualRouterDelete.json
     */
    /**
     * Sample code: Delete VirtualRouter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteVirtualRouter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualRouters().delete("rg1", "virtualRouter", Context.NONE);
    }
}
