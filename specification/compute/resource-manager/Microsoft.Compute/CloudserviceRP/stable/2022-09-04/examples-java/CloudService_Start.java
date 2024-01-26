
/**
 * Samples for CloudServices Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudService_Start.json
     */
    /**
     * Sample code: Start Cloud Service.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startCloudService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServices().start("ConstosoRG", "{cs-name}",
            com.azure.core.util.Context.NONE);
    }
}
