
/**
 * Samples for ServerInstances ListBySapInstance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/
     * examples/ServerInstances_List.json
     */
    /**
     * Sample code: Lists the Server Instance resources for the given SAP Instance resource.
     * 
     * @param manager Entry point to MigrationDiscoverySapManager.
     */
    public static void listsTheServerInstanceResourcesForTheGivenSAPInstanceResource(
        com.azure.resourcemanager.migrationdiscoverysap.MigrationDiscoverySapManager manager) {
        manager.serverInstances().listBySapInstance("test-rg", "SampleSite", "MPP_MPP",
            com.azure.core.util.Context.NONE);
    }
}
