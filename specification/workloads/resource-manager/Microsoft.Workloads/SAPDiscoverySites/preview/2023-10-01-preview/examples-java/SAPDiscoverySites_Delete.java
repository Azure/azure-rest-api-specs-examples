
/**
 * Samples for SapDiscoverySites Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/SAPDiscoverySites_Delete.json
     */
    /**
     * Sample code: Deletes a SAP Migration discovery site resource.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void deletesASAPMigrationDiscoverySiteResource(
        com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        manager.sapDiscoverySites().delete("test-rg", "SampleSite", com.azure.core.util.Context.NONE);
    }
}
