import com.azure.core.util.Context;

/** Samples for AvailableDelegations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/AvailableDelegationsSubscriptionGet.json
     */
    /**
     * Sample code: Get available delegations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableDelegations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAvailableDelegations().list("westcentralus", Context.NONE);
    }
}
