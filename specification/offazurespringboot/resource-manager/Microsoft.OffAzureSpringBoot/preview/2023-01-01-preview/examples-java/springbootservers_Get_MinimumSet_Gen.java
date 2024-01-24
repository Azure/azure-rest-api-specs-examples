
/**
 * Samples for Springbootservers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootservers_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: springbootservers_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootserversGetMinimumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootservers().getWithResponse("rgspringbootservers",
            "hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj", "zkarbqnwnxeozvjrkpdqmgnwedwgtwcmmyqwaijkn",
            com.azure.core.util.Context.NONE);
    }
}
