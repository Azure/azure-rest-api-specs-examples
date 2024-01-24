
/**
 * Samples for Springbootsites List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootsites_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: springbootsites_ListBySubscription_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootsitesListBySubscriptionMinimumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootsites().list(com.azure.core.util.Context.NONE);
    }
}
