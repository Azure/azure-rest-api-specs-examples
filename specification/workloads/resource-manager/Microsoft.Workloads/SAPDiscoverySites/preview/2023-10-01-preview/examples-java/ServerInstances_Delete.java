
/**
 * Samples for ServerInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/ServerInstances_Delete.json
     */
    /**
     * Sample code: Deletes the Server Instance resource.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void deletesTheServerInstanceResource(
        com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        manager.serverInstances().delete("test-rg", "SampleSite", "MPP_MPP", "APP_SapServer1",
            com.azure.core.util.Context.NONE);
    }
}
