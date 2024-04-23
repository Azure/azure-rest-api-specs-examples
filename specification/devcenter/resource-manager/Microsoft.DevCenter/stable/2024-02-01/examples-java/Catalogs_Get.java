
/**
 * Samples for Catalogs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Catalogs_Get.json
     */
    /**
     * Sample code: Catalogs_Get.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.catalogs().getWithResponse("rg1", "Contoso", "CentralCatalog", com.azure.core.util.Context.NONE);
    }
}
