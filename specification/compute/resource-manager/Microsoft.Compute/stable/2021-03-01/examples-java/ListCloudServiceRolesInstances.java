import com.azure.core.util.Context;

/** Samples for CloudServiceRoleInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListCloudServiceRolesInstances.json
     */
    /**
     * Sample code: List Role Instances in a Cloud Service.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRoleInstancesInACloudService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServiceRoleInstances()
            .list("ConstosoRG", "{cs-name}", null, Context.NONE);
    }
}
