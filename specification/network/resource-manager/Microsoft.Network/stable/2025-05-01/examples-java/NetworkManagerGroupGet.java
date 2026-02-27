
/**
 * Samples for NetworkGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerGroupGet.json
     */
    /**
     * Sample code: NetworkGroupsGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkGroupsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkGroups().getWithResponse("rg1", "testNetworkManager",
            "testNetworkGroup", com.azure.core.util.Context.NONE);
    }
}
