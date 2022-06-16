import com.azure.core.util.Context;

/** Samples for CloudServices GetInstanceView. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetCloudServiceInstanceViewWithMultiRole.json
     */
    /**
     * Sample code: Get Cloud Service Instance View with Multiple Roles.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceInstanceViewWithMultipleRoles(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServices()
            .getInstanceViewWithResponse("ConstosoRG", "{cs-name}", Context.NONE);
    }
}
