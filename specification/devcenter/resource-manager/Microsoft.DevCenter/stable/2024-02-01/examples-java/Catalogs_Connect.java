
/**
 * Samples for Catalogs Connect.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Catalogs_Connect.json
     */
    /**
     * Sample code: Catalogs_Connect.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsConnect(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.catalogs().connect("rg1", "Contoso", "CentralCatalog", com.azure.core.util.Context.NONE);
    }
}
