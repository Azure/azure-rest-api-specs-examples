
/**
 * Samples for SapInstances ListBySapDiscoverySite.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/SAPInstances_List.json
     */
    /**
     * Sample code: Lists the SAP Instance resources for the given SAP Migration discovery site resource.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void listsTheSAPInstanceResourcesForTheGivenSAPMigrationDiscoverySiteResource(
        com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        manager.sapInstances().listBySapDiscoverySite("test-rg", "SampleSite", com.azure.core.util.Context.NONE);
    }
}
