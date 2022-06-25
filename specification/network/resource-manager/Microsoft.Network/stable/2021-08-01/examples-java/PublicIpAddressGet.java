import com.azure.core.util.Context;

/** Samples for PublicIpAddresses GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PublicIpAddressGet.json
     */
    /**
     * Sample code: Get public IP address.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPublicIPAddress(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPublicIpAddresses()
            .getByResourceGroupWithResponse("rg1", "testDNS-ip", null, Context.NONE);
    }
}
