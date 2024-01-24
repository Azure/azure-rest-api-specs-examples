
/**
 * Samples for Springbootsites GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootsites_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: springbootsites_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootsitesGetMaximumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootsites().getByResourceGroupWithResponse("rgspringbootsites", "xrmzlavpewxtfeitghdrj",
            com.azure.core.util.Context.NONE);
    }
}
