
/**
 * Samples for CloudServicesUpdateDomain GetUpdateDomain.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceUpdateDomain_Get.json
     */
    /**
     * Sample code: Get Cloud Service Update Domain.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceUpdateDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServicesUpdateDomains()
            .getUpdateDomainWithResponse("ConstosoRG", "{cs-name}", 1, com.azure.core.util.Context.NONE);
    }
}
