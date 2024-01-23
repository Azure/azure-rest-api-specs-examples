
/**
 * Samples for Springbootapps ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootapps_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: springbootapps_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootappsListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootapps().listBySubscription("pdfosfhtemfsaglvwjdyqlyeipucrd", com.azure.core.util.Context.NONE);
    }
}
