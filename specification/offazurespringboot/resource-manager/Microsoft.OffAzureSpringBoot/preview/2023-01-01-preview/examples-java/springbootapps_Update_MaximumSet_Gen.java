
import com.azure.resourcemanager.springappdiscovery.models.SpringbootappsPatch;

/**
 * Samples for Springbootapps Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootapps_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: springbootapps_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootappsUpdateMaximumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootapps().update("rgspringbootapps", "pdfosfhtemfsaglvwjdyqlyeipucrd",
            "ofjeesoahqtnovlbuvflyknpbhcpeqqhekntvqxyemuwbcqnuxjgfhsf", new SpringbootappsPatch(),
            com.azure.core.util.Context.NONE);
    }
}
