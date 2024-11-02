
/**
 * Samples for ErrorSummaries Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/ErrorSummaries_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: ErrorSummaries_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void
        errorSummariesGetMinimumSetGen(com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.errorSummaries().getWithResponse("rgspringbootdiscovery", "xxkzlvbihwxunadjcpjpjmghmhxrqyvghtpfps",
            "K2lv", com.azure.core.util.Context.NONE);
    }
}
