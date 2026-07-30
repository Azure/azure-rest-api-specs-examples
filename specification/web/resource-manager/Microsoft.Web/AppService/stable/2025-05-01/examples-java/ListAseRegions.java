
/**
 * Samples for ResourceProviders ListAseRegions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListAseRegions.json
     */
    /**
     * Sample code: List aseregions.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAseregions(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getResourceProviders().listAseRegions(com.azure.core.util.Context.NONE);
    }
}
