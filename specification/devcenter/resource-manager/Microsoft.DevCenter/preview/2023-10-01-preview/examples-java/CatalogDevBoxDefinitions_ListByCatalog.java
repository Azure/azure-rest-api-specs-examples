/** Samples for CatalogDevBoxDefinitions ListByCatalog. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/CatalogDevBoxDefinitions_ListByCatalog.json
     */
    /**
     * Sample code: CatalogDevBoxDefinitions_ListByCatalog.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogDevBoxDefinitionsListByCatalog(
        com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .catalogDevBoxDefinitions()
            .listByCatalog("rg1", "Contoso", "CentralCatalog", null, com.azure.core.util.Context.NONE);
    }
}
