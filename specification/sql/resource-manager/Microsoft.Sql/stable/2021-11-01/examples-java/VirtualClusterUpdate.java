
import com.azure.resourcemanager.sql.models.VirtualClusterUpdate;

/**
 * Samples for VirtualClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/VirtualClusterUpdate.json
     */
    /**
     * Sample code: Update virtual cluster with tags.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateVirtualClusterWithTags(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getVirtualClusters().update("testrg",
            "vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2",
            new VirtualClusterUpdate().withMaintenanceConfigurationId(
                "/subscriptions/ab0e51c0-83c0-4380-8ae9-025516df392f/resourceGroups/Federation/providers/Microsoft.Maintenance/maintenanceConfigurations/MiPolicy1"),
            com.azure.core.util.Context.NONE);
    }
}
