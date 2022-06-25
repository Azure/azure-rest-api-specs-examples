import com.azure.core.util.Context;

/** Samples for VirtualHubRouteTableV2S Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualHubRouteTableV2Get.json
     */
    /**
     * Sample code: VirtualHubVirtualHubRouteTableV2Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubVirtualHubRouteTableV2Get(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualHubRouteTableV2S()
            .getWithResponse("rg1", "virtualHub1", "virtualHubRouteTable1a", Context.NONE);
    }
}
