import com.azure.core.util.Context;

/** Samples for CloudServiceRoles Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetCloudServiceRole.json
     */
    /**
     * Sample code: Get Cloud Service Role.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceRole(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServiceRoles()
            .getWithResponse("{role-name}", "ConstosoRG", "{cs-name}", Context.NONE);
    }
}
