import com.azure.core.util.Context;

/** Samples for CloudServiceRoleInstances Reimage. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ReimageCloudServiceRoleInstance.json
     */
    /**
     * Sample code: Reimage Cloud Service Role Instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reimageCloudServiceRoleInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServiceRoleInstances()
            .reimage("{roleInstance-name}", "ConstosoRG", "{cs-name}", Context.NONE);
    }
}
