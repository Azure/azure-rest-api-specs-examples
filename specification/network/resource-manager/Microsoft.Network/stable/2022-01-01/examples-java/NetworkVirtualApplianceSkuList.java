import com.azure.core.util.Context;

/** Samples for VirtualApplianceSkus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkVirtualApplianceSkuList.json
     */
    /**
     * Sample code: NetworkVirtualApplianceSkuListResult.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkVirtualApplianceSkuListResult(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualApplianceSkus().list(Context.NONE);
    }
}
