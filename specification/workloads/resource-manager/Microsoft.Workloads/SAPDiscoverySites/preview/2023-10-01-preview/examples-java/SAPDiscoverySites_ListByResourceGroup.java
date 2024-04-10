
/**
 * Samples for SapDiscoverySites ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/SAPDiscoverySites_ListByResourceGroup.json
     */
    /**
     * Sample code: List SAP Migration discovery site resources by Resource group.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void listSAPMigrationDiscoverySiteResourcesByResourceGroup(
        com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        manager.sapDiscoverySites().listByResourceGroup("test-rg", com.azure.core.util.Context.NONE);
    }
}
