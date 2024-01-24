
/**
 * Samples for Springbootsites ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootsites_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: springbootsites_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootsitesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootsites().listByResourceGroup("rgspringbootsites", com.azure.core.util.Context.NONE);
    }
}
