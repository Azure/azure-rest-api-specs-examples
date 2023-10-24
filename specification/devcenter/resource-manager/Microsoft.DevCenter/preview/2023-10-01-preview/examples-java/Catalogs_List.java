/** Samples for Catalogs ListByDevCenter. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Catalogs_List.json
     */
    /**
     * Sample code: Catalogs_ListByDevCenter.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsListByDevCenter(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.catalogs().listByDevCenter("rg1", "Contoso", null, com.azure.core.util.Context.NONE);
    }
}
