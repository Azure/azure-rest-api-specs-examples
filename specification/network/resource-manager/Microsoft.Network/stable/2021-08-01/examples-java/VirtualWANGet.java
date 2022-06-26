import com.azure.core.util.Context;

/** Samples for VirtualWans GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualWANGet.json
     */
    /**
     * Sample code: VirtualWANGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualWANGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualWans()
            .getByResourceGroupWithResponse("rg1", "wan1", Context.NONE);
    }
}
