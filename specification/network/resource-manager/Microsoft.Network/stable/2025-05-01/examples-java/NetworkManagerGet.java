
/**
 * Samples for NetworkManagers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerGet.json
     */
    /**
     * Sample code: NetworkManagersGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkManagersGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkManagers().getByResourceGroupWithResponse("rg1",
            "testNetworkManager", com.azure.core.util.Context.NONE);
    }
}
