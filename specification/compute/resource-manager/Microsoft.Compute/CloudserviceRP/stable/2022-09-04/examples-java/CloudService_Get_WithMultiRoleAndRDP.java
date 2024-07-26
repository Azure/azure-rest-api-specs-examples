
/**
 * Samples for CloudServices GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudService_Get_WithMultiRoleAndRDP.json
     */
    /**
     * Sample code: Get Cloud Service with Multiple Roles and RDP Extension.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getCloudServiceWithMultipleRolesAndRDPExtension(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServices()
            .getByResourceGroupWithResponse("ConstosoRG", "{cs-name}", com.azure.core.util.Context.NONE);
    }
}
