
/**
 * Samples for CloudServices GetInstanceView.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudService_Get_InstanceViewWithMultiRole.json
     */
    /**
     * Sample code: Get Cloud Service Instance View with Multiple Roles.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getCloudServiceInstanceViewWithMultipleRoles(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServices().getInstanceViewWithResponse("ConstosoRG",
            "{cs-name}", com.azure.core.util.Context.NONE);
    }
}
