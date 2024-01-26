
/**
 * Samples for CloudServices List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudService_List_BySubscription.json
     */
    /**
     * Sample code: List Cloud Services in a Subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCloudServicesInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServices().list(com.azure.core.util.Context.NONE);
    }
}
