import com.azure.core.util.Context;

/** Samples for PublicIpPrefixes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/PublicIpPrefixListAll.json
     */
    /**
     * Sample code: List all public IP prefixes.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllPublicIPPrefixes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpPrefixes().list(Context.NONE);
    }
}
