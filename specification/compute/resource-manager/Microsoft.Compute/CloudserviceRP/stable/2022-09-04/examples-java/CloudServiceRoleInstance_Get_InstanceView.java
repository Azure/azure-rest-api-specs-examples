
/**
 * Samples for CloudServiceRoleInstances GetInstanceView.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceRoleInstance_Get_InstanceView.json
     */
    /**
     * Sample code: Get Instance View of Cloud Service Role Instance.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInstanceViewOfCloudServiceRoleInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServiceRoleInstances().getInstanceViewWithResponse(
            "{roleInstance-name}", "ConstosoRG", "{cs-name}", com.azure.core.util.Context.NONE);
    }
}
