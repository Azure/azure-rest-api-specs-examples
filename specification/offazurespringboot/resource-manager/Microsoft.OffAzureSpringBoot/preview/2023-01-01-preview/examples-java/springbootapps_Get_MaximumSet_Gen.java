
/**
 * Samples for Springbootapps Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootapps_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: springbootapps_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void
        springbootappsGetMaximumSetGen(com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootapps().getWithResponse("rgspringbootapps", "pdfosfhtemfsaglvwjdyqlyeipucrd",
            "ofjeesoahqtnovlbuvflyknpbhcpeqqhekntvqxyemuwbcqnuxjgfhsf", com.azure.core.util.Context.NONE);
    }
}
