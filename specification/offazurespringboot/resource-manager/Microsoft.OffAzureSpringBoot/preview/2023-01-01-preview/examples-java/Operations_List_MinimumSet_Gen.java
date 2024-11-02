
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void
        operationsListMinimumSetGen(com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
