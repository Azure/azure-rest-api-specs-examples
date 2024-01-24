
/**
 * Samples for Springbootservers ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootservers_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: springbootservers_ListBySubscription_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootserversListBySubscriptionMinimumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootservers().listBySubscription("hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj",
            com.azure.core.util.Context.NONE);
    }
}
