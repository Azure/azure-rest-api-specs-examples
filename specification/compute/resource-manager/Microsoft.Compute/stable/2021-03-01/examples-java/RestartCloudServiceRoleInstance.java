import com.azure.core.util.Context;

/** Samples for CloudServiceRoleInstances Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/RestartCloudServiceRoleInstance.json
     */
    /**
     * Sample code: Restart Cloud Service Role Instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restartCloudServiceRoleInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServiceRoleInstances()
            .restart("{roleInstance-name}", "ConstosoRG", "{cs-name}", Context.NONE);
    }
}
