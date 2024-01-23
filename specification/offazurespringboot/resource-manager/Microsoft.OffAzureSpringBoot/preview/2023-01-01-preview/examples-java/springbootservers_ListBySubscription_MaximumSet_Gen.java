
/**
 * Samples for Springbootservers ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootservers_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: springbootservers_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootserversListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootservers().listBySubscription("hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj",
            com.azure.core.util.Context.NONE);
    }
}
