
/**
 * Samples for Springbootservers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootservers_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: springbootservers_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootserversDeleteMaximumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootservers().delete("rgspringbootservers",
            "hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj", "zkarbqnwnxeozvjrkpdqmgnwedwgtwcmmyqwaijkn",
            com.azure.core.util.Context.NONE);
    }
}
