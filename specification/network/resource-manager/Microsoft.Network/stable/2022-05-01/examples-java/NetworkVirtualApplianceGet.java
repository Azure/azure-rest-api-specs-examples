import com.azure.core.util.Context;

/** Samples for NetworkVirtualAppliances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkVirtualApplianceGet.json
     */
    /**
     * Sample code: Get NetworkVirtualAppliance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getNetworkVirtualAppliance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkVirtualAppliances()
            .getByResourceGroupWithResponse("rg1", "nva", null, Context.NONE);
    }
}
