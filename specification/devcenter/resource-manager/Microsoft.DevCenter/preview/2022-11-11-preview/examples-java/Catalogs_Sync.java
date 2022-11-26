import com.azure.core.util.Context;

/** Samples for Catalogs Sync. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Catalogs_Sync.json
     */
    /**
     * Sample code: Catalogs_Sync.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsSync(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.catalogs().sync("rg1", "Contoso", "CentralCatalog", Context.NONE);
    }
}
