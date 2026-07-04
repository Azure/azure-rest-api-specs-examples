
/**
 * Samples for BastionHosts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BastionHostListBySubscription.json
     */
    /**
     * Sample code: List all Bastion Hosts for a given subscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllBastionHostsForAGivenSubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().list(com.azure.core.util.Context.NONE);
    }
}
