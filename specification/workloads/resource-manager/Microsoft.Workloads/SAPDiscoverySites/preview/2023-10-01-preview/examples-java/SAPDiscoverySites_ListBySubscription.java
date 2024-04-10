
/**
 * Samples for SapDiscoverySites List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/SAPDiscoverySites_ListBySubscription.json
     */
    /**
     * Sample code: List SAP Migration discovery site resources in a subscription.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void listSAPMigrationDiscoverySiteResourcesInASubscription(
        com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        manager.sapDiscoverySites().list(com.azure.core.util.Context.NONE);
    }
}
