
import com.azure.resourcemanager.network.fluent.models.PublicIpAddressInner;

/**
 * Samples for PublicIpAddresses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * PublicIpAddressCreateDefaultsStandardV2Sku.json
     */
    /**
     * Sample code: Create public IP address defaults with StandardV2 Sku.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createPublicIPAddressDefaultsWithStandardV2Sku(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().createOrUpdate("rg1", "test-ip",
            new PublicIpAddressInner().withLocation("eastus"), com.azure.core.util.Context.NONE);
    }
}
