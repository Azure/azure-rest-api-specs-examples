
/**
 * Samples for CloudServiceOperatingSystems ListOSFamilies.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceOSFamilies_List.json
     */
    /**
     * Sample code: List Cloud Service OS Families in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCloudServiceOSFamiliesInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServiceOperatingSystems().listOSFamilies("westus2",
            com.azure.core.util.Context.NONE);
    }
}
