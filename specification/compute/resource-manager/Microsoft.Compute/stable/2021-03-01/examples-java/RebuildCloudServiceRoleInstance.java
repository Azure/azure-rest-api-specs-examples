import com.azure.core.util.Context;

/** Samples for CloudServiceRoleInstances Rebuild. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/RebuildCloudServiceRoleInstance.json
     */
    /**
     * Sample code: Rebuild Cloud Service Role Instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void rebuildCloudServiceRoleInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServiceRoleInstances()
            .rebuild("{roleInstance-name}", "ConstosoRG", "{cs-name}", Context.NONE);
    }
}
