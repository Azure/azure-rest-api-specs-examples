
/**
 * Samples for CloudServices PowerOff.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudService_PowerOff.json
     */
    /**
     * Sample code: Stop or PowerOff Cloud Service.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void stopOrPowerOffCloudService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServices().powerOff("ConstosoRG", "{cs-name}",
            com.azure.core.util.Context.NONE);
    }
}
