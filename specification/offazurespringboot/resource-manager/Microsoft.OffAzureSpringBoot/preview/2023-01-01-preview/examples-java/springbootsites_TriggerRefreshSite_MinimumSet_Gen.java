
/**
 * Samples for Springbootsites TriggerRefreshSite.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootsites_TriggerRefreshSite_MinimumSet_Gen.json
     */
    /**
     * Sample code: springbootsites_TriggerRefreshSite_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootsitesTriggerRefreshSiteMinimumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootsites().triggerRefreshSite("rgspringbootsites", "czarpuxwoafaqsuptutcwyu",
            com.azure.core.util.Context.NONE);
    }
}
