import com.azure.core.util.Context;

/** Samples for NetworkVirtualAppliances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkVirtualApplianceListBySubscription.json
     */
    /**
     * Sample code: List all Network Virtual Appliances for a given subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllNetworkVirtualAppliancesForAGivenSubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkVirtualAppliances().list(Context.NONE);
    }
}
