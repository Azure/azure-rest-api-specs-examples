
/**
 * Samples for CustomIpPrefixes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CustomIpPrefixDelete.json
     */
    /**
     * Sample code: Delete custom IP prefix.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteCustomIPPrefix(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getCustomIpPrefixes().delete("rg1", "test-customipprefix",
            com.azure.core.util.Context.NONE);
    }
}
