import com.azure.core.util.Context;

/** Samples for CloudServiceRoles List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListCloudServiceRoles.json
     */
    /**
     * Sample code: List Roles in a Cloud Service.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRolesInACloudService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServiceRoles()
            .list("ConstosoRG", "{cs-name}", Context.NONE);
    }
}
