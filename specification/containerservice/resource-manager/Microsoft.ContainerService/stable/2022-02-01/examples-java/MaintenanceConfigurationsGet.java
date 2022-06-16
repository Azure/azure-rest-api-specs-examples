import com.azure.core.util.Context;

/** Samples for MaintenanceConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-02-01/examples/MaintenanceConfigurationsGet.json
     */
    /**
     * Sample code: Get Maintenance Configuration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getMaintenanceConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getMaintenanceConfigurations()
            .getWithResponse("rg1", "clustername1", "default", Context.NONE);
    }
}
