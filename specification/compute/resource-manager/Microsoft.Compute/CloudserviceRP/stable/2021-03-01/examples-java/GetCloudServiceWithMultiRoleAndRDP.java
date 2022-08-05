import com.azure.core.util.Context;

/** Samples for CloudServices GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetCloudServiceWithMultiRoleAndRDP.json
     */
    /**
     * Sample code: Get Cloud Service with Multiple Roles and RDP Extension.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceWithMultipleRolesAndRDPExtension(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServices()
            .getByResourceGroupWithResponse("ConstosoRG", "{cs-name}", Context.NONE);
    }
}
