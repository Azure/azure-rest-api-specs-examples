import com.azure.core.util.Context;

/** Samples for MaintenanceConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-10-01/examples/MaintenanceConfigurationsDelete.json
     */
    /**
     * Sample code: Delete Maintenance Configuration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteMaintenanceConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getMaintenanceConfigurations()
            .deleteWithResponse("rg1", "clustername1", "default", Context.NONE);
    }
}
