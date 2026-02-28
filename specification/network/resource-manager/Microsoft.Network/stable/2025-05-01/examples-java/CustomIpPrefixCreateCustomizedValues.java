
import com.azure.resourcemanager.network.fluent.models.CustomIpPrefixInner;

/**
 * Samples for CustomIpPrefixes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * CustomIpPrefixCreateCustomizedValues.json
     */
    /**
     * Sample code: Create custom IP prefix allocation method.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createCustomIPPrefixAllocationMethod(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getCustomIpPrefixes().createOrUpdate("rg1", "test-customipprefix",
            new CustomIpPrefixInner().withLocation("westus").withCidr("0.0.0.0/24"), com.azure.core.util.Context.NONE);
    }
}
