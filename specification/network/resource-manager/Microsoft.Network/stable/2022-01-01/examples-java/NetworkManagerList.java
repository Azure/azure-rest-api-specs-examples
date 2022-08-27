import com.azure.core.util.Context;

/** Samples for NetworkManagers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerList.json
     */
    /**
     * Sample code: List Network Manager.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNetworkManager(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkManagers()
            .listByResourceGroup("rg1", null, null, Context.NONE);
    }
}
