
/**
 * Samples for Springbootsites Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootsites_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: springbootsites_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootsitesDeleteMinimumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootsites().delete("rgspringbootsites", "xrmzlavpewxtfeitghdrj",
            com.azure.core.util.Context.NONE);
    }
}
