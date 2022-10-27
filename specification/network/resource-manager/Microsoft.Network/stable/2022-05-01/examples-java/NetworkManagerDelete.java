import com.azure.core.util.Context;

/** Samples for NetworkManagers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkManagerDelete.json
     */
    /**
     * Sample code: NetworkManagersDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkManagersDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkManagers()
            .delete("rg1", "testNetworkManager", false, Context.NONE);
    }
}
