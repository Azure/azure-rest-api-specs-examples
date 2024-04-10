
/**
 * Samples for ServerInstances Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/ServerInstances_Create.json
     */
    /**
     * Sample code: Creates the Server Instance resource.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void createsTheServerInstanceResource(
        com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        manager.serverInstances().define("APP_SapServer1").withExistingSapInstance("test-rg", "SampleSite", "MPP_MPP")
            .create();
    }
}
