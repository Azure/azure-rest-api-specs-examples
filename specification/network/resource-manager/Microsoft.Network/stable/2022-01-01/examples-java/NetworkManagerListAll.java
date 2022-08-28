import com.azure.core.util.Context;

/** Samples for NetworkManagers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerListAll.json
     */
    /**
     * Sample code: NetworkManagersList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkManagersList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkManagers().list(null, null, Context.NONE);
    }
}
