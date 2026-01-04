
import com.azure.resourcemanager.network.fluent.models.PublicIpAddressInner;

/**
 * Samples for PublicIpAddresses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PublicIpAddressCreateDefaults
     * .json
     */
    /**
     * Sample code: Create public IP address defaults.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createPublicIPAddressDefaults(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().createOrUpdate("rg1", "test-ip",
            new PublicIpAddressInner().withLocation("eastus"), com.azure.core.util.Context.NONE);
    }
}
