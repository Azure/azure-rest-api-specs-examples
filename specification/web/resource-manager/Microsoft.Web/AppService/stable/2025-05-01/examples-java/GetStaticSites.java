
/**
 * Samples for StaticSites ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSites.json
     */
    /**
     * Sample code: Get static sites for a resource group.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getStaticSitesForAResourceGroup(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
