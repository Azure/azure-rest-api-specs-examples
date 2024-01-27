
/**
 * Samples for CloudServiceRoles Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceRole_Get.json
     */
    /**
     * Sample code: Get Cloud Service Role.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceRole(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServiceRoles().getWithResponse("{role-name}",
            "ConstosoRG", "{cs-name}", com.azure.core.util.Context.NONE);
    }
}
