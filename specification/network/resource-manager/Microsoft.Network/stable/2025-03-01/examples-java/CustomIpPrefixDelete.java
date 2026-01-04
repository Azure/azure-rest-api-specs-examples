
/**
 * Samples for CustomIpPrefixes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/CustomIpPrefixDelete.json
     */
    /**
     * Sample code: Delete custom IP prefix.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteCustomIPPrefix(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getCustomIpPrefixes().delete("rg1", "test-customipprefix",
            com.azure.core.util.Context.NONE);
    }
}
