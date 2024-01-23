
import com.azure.resourcemanager.springappdiscovery.models.SpringbootserversModel;

/**
 * Samples for Springbootservers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootservers_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: springbootservers_Update_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootserversUpdateMinimumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        SpringbootserversModel resource = manager.springbootservers()
            .getWithResponse("rgspringbootservers", "hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj",
                "zkarbqnwnxeozvjrkpdqmgnwedwgtwcmmyqwaijkn", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }
}
