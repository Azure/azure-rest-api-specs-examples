
/**
 * Samples for CloudServiceRoleInstances GetRemoteDesktopFile.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceRoleInstance_Get_RemoteDesktopFile.json
     */
    /**
     * Sample code: Get Cloud Service Role.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCloudServiceRole(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServiceRoleInstances()
            .getRemoteDesktopFileWithResponse("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "rgcloudService", "aaaa",
                com.azure.core.util.Context.NONE);
    }
}
