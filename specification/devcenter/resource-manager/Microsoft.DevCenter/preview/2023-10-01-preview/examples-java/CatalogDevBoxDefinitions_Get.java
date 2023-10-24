/** Samples for CatalogDevBoxDefinitions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/CatalogDevBoxDefinitions_Get.json
     */
    /**
     * Sample code: CatalogDevBoxDefinitions_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogDevBoxDefinitionsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .catalogDevBoxDefinitions()
            .getWithResponse("rg1", "Contoso", "CentralCatalog", "WebDevBox", com.azure.core.util.Context.NONE);
    }
}
