
/**
 * Samples for PublicIpPrefixes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpPrefixDelete.json
     */
    /**
     * Sample code: Delete public IP prefix.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletePublicIPPrefix(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpPrefixes().delete("rg1", "test-ipprefix", com.azure.core.util.Context.NONE);
    }
}
