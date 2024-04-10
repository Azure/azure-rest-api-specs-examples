
/**
 * Samples for SapInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/SAPInstances_Delete.json
     */
    /**
     * Sample code: Deletes the SAP Instance resource.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void deletesTheSAPInstanceResource(
        com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        manager.sapInstances().delete("test-rg", "SampleSite", "MPP_MPP", com.azure.core.util.Context.NONE);
    }
}
