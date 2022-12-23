import com.azure.core.util.Context;

/** Samples for PublicIpAddresses Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/PublicIpAddressDelete.json
     */
    /**
     * Sample code: Delete public IP address.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletePublicIPAddress(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().delete("rg1", "test-ip", Context.NONE);
    }
}
