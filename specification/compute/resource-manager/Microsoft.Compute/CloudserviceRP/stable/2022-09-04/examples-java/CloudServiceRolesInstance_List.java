
/**
 * Samples for CloudServiceRoleInstances List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceRolesInstance_List.json
     */
    /**
     * Sample code: List Role Instances in a Cloud Service.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRoleInstancesInACloudService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServiceRoleInstances().list("ConstosoRG", "{cs-name}",
            null, com.azure.core.util.Context.NONE);
    }
}
