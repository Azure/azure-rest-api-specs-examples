import com.azure.core.util.Context;

/** Samples for CloudServiceRoleInstances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetCloudServiceRoleInstance.json
     */
    /**
     * Sample code: Get Cloud Service Role Instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceRoleInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServiceRoleInstances()
            .getWithResponse("{roleInstance-name}", "ConstosoRG", "{cs-name}", null, Context.NONE);
    }
}
