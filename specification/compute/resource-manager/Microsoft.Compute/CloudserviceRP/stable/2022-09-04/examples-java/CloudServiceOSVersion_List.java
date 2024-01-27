
/**
 * Samples for CloudServiceOperatingSystems ListOSVersions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceOSVersion_List.json
     */
    /**
     * Sample code: List Cloud Service OS Versions in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCloudServiceOSVersionsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServiceOperatingSystems().listOSVersions("westus2",
            com.azure.core.util.Context.NONE);
    }
}
