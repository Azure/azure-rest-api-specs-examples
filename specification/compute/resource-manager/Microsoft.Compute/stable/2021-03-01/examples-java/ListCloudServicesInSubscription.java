import com.azure.core.util.Context;

/** Samples for CloudServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListCloudServicesInSubscription.json
     */
    /**
     * Sample code: List Cloud Services in a Subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCloudServicesInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServices().list(Context.NONE);
    }
}
