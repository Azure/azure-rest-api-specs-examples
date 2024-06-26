
import com.azure.resourcemanager.migrationdiscoverysap.models.ServerInstance;
import com.azure.resourcemanager.migrationdiscoverysap.models.ServerInstanceProperties;

/**
 * Samples for ServerInstances Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/ServerInstances_Update.json
     */
    /**
     * Sample code: Updates the Server Instance resource.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void updatesTheServerInstanceResource(
        com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        ServerInstance resource = manager.serverInstances()
            .getWithResponse("test-rg", "SampleSite", "MPP_MPP", "APP_SapServer1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new ServerInstanceProperties()).apply();
    }
}
