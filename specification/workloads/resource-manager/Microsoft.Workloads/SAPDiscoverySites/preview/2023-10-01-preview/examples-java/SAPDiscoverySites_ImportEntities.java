
/**
 * Samples for SapDiscoverySites ImportEntities.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/SAPDiscoverySites_ImportEntities.json
     */
    /**
     * Sample code: Import a SAP Migration discovery site resource and it's child resources.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void importASAPMigrationDiscoverySiteResourceAndItSChildResources(
        com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        manager.sapDiscoverySites().importEntities("test-rg", "SampleSite", com.azure.core.util.Context.NONE);
    }
}
