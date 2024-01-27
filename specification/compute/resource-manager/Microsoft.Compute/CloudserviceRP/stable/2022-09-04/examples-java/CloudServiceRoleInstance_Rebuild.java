
/**
 * Samples for CloudServiceRoleInstances Rebuild.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceRoleInstance_Rebuild.json
     */
    /**
     * Sample code: Rebuild Cloud Service Role Instance.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void rebuildCloudServiceRoleInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServiceRoleInstances().rebuild("{roleInstance-name}",
            "ConstosoRG", "{cs-name}", com.azure.core.util.Context.NONE);
    }
}
