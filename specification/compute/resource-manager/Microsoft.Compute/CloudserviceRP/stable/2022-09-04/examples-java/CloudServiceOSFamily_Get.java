
/**
 * Samples for CloudServiceOperatingSystems GetOSFamily.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceOSFamily_Get.json
     */
    /**
     * Sample code: Get Cloud Service OS Family.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceOSFamily(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServiceOperatingSystems()
            .getOSFamilyWithResponse("westus2", "3", com.azure.core.util.Context.NONE);
    }
}
