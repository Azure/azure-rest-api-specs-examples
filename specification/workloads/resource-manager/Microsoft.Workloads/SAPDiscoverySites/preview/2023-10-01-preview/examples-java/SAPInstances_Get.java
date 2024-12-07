
/**
 * Samples for SapInstances Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/SAPInstances_Get.json
     */
    /**
     * Sample code: GET a SAP Instance resource.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void
        gETASAPInstanceResource(com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        manager.sapInstances().getWithResponse("test-rg", "SampleSite", "MPP_MPP", com.azure.core.util.Context.NONE);
    }
}
