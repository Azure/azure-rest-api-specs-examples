import com.azure.core.util.Context;

/** Samples for PublicIpPrefixes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PublicIpPrefixDelete.json
     */
    /**
     * Sample code: Delete public IP prefix.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletePublicIPPrefix(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpPrefixes().delete("rg1", "test-ipprefix", Context.NONE);
    }
}
