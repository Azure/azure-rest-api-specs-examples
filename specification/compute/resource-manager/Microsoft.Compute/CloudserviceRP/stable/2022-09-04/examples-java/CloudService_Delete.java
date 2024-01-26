
/**
 * Samples for CloudServices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudService_Delete.json
     */
    /**
     * Sample code: Delete Cloud Service.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteCloudService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServices().delete("ConstosoRG", "{cs-name}",
            com.azure.core.util.Context.NONE);
    }
}
