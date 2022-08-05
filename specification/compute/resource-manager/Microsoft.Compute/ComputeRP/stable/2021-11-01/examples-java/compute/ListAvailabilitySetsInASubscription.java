import com.azure.core.util.Context;

/** Samples for AvailabilitySets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/ListAvailabilitySetsInASubscription.json
     */
    /**
     * Sample code: List availability sets in a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAvailabilitySetsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getAvailabilitySets()
            .list("virtualMachines\\$ref", Context.NONE);
    }
}
