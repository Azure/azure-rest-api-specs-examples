
/**
 * Samples for Summaries Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/Summaries_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: Summaries_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void
        summariesGetMinimumSetGen(com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.summaries().getWithResponse("rgspringbootdiscovery", "xxkzlvbihwxunadjcpjpjmghmhxrqyvghtpfps", "vjB",
            com.azure.core.util.Context.NONE);
    }
}
