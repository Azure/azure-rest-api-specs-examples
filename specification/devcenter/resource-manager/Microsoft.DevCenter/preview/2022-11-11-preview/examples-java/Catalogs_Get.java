import com.azure.core.util.Context;

/** Samples for Catalogs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Catalogs_Get.json
     */
    /**
     * Sample code: Catalogs_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.catalogs().getWithResponse("rg1", "Contoso", "CentralCatalog", Context.NONE);
    }
}
