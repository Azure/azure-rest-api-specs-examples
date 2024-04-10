
/**
 * Samples for ServerInstances Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/ServerInstances_Get.json
     */
    /**
     * Sample code: GET a Server Instance resource.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void gETAServerInstanceResource(
        com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        manager.serverInstances().getWithResponse("test-rg", "SampleSite", "MPP_MPP", "APP_SapServer1",
            com.azure.core.util.Context.NONE);
    }
}
