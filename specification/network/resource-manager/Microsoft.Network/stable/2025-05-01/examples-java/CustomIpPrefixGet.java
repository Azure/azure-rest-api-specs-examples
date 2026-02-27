
/**
 * Samples for CustomIpPrefixes GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/CustomIpPrefixGet.json
     */
    /**
     * Sample code: Get custom IP prefix.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCustomIPPrefix(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getCustomIpPrefixes().getByResourceGroupWithResponse("rg1",
            "test-customipprefix", null, com.azure.core.util.Context.NONE);
    }
}
