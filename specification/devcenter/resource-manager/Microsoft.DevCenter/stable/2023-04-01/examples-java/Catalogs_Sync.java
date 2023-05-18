/** Samples for Catalogs Sync. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Catalogs_Sync.json
     */
    /**
     * Sample code: Catalogs_Sync.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsSync(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.catalogs().sync("rg1", "Contoso", "CentralCatalog", com.azure.core.util.Context.NONE);
    }
}
