
/**
 * Samples for Springbootservers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootservers_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: springbootservers_CreateOrUpdate_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootserversCreateOrUpdateMinimumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootservers().define("zkarbqnwnxeozvjrkpdqmgnwedwgtwcmmyqwaijkn").withExistingSpringbootsite(
            "rgspringbootservers", "hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj").create();
    }
}
