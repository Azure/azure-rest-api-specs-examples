/** Samples for Catalogs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Catalogs_Delete.json
     */
    /**
     * Sample code: Catalogs_Delete.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.catalogs().delete("rg1", "Contoso", "CentralCatalog", com.azure.core.util.Context.NONE);
    }
}
