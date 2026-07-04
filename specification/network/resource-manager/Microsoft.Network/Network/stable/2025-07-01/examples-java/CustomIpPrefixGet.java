
/**
 * Samples for CustomIpPrefixes GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CustomIpPrefixGet.json
     */
    /**
     * Sample code: Get custom IP prefix.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getCustomIPPrefix(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getCustomIpPrefixes().getByResourceGroupWithResponse("rg1", "test-customipprefix", null,
            com.azure.core.util.Context.NONE);
    }
}
