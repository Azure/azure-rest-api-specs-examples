
/**
 * Samples for CloudServiceRoles List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceRole_List.json
     */
    /**
     * Sample code: List Roles in a Cloud Service.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRolesInACloudService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServiceRoles().list("ConstosoRG", "{cs-name}",
            com.azure.core.util.Context.NONE);
    }
}
