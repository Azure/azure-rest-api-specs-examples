
/**
 * Samples for Summaries ListBySite.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/Summaries_ListBySite_MinimumSet_Gen.json
     */
    /**
     * Sample code: Summaries_ListBySite_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void summariesListBySiteMinimumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.summaries().listBySite("rgspringbootdiscovery", "xxkzlvbihwxunadjcpjpjmghmhxrqyvghtpfps",
            com.azure.core.util.Context.NONE);
    }
}
