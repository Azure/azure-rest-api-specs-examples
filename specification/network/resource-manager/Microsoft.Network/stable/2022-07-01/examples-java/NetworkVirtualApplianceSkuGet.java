import com.azure.core.util.Context;

/** Samples for VirtualApplianceSkus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkVirtualApplianceSkuGet.json
     */
    /**
     * Sample code: NetworkVirtualApplianceSkuGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkVirtualApplianceSkuGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualApplianceSkus()
            .getWithResponse("ciscoSdwan", Context.NONE);
    }
}
