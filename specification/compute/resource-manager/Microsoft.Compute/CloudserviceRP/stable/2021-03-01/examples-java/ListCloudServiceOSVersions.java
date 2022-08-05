import com.azure.core.util.Context;

/** Samples for CloudServiceOperatingSystems ListOSVersions. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListCloudServiceOSVersions.json
     */
    /**
     * Sample code: List Cloud Service OS Versions in a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCloudServiceOSVersionsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServiceOperatingSystems()
            .listOSVersions("westus2", Context.NONE);
    }
}
