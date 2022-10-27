import com.azure.core.util.Context;

/** Samples for BastionHosts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/BastionHostListBySubscription.json
     */
    /**
     * Sample code: List all Bastion Hosts for a given subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllBastionHostsForAGivenSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getBastionHosts().list(Context.NONE);
    }
}
