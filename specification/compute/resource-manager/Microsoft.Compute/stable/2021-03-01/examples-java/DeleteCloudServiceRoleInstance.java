import com.azure.core.util.Context;

/** Samples for CloudServiceRoleInstances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/DeleteCloudServiceRoleInstance.json
     */
    /**
     * Sample code: Delete Cloud Service Role Instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteCloudServiceRoleInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServiceRoleInstances()
            .delete("{roleInstance-name}", "ConstosoRG", "{cs-name}", Context.NONE);
    }
}
