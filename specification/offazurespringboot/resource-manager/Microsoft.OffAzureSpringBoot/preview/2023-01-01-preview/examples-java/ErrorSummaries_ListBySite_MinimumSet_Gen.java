
/**
 * Samples for ErrorSummaries ListBySite.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/ErrorSummaries_ListBySite_MinimumSet_Gen.json
     */
    /**
     * Sample code: ErrorSummaries_ListBySite_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void errorSummariesListBySiteMinimumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.errorSummaries().listBySite("rgspringbootdiscovery", "xxkzlvbihwxunadjcpjpjmghmhxrqyvghtpfps",
            com.azure.core.util.Context.NONE);
    }
}
