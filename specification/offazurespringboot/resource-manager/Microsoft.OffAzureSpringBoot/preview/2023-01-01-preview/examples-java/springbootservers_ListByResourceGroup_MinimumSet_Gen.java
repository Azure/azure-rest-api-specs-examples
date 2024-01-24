
/**
 * Samples for Springbootservers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootservers_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: springbootservers_ListByResourceGroup_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootserversListByResourceGroupMinimumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootservers().listByResourceGroup("rgspringbootservers",
            "hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj", com.azure.core.util.Context.NONE);
    }
}
