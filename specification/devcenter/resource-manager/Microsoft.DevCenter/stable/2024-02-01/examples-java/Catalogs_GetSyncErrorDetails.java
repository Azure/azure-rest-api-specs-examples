
/**
 * Samples for Catalogs GetSyncErrorDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * Catalogs_GetSyncErrorDetails.json
     */
    /**
     * Sample code: Catalogs_GetSyncErrorDetails.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsGetSyncErrorDetails(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.catalogs().getSyncErrorDetailsWithResponse("rg1", "Contoso", "CentralCatalog",
            com.azure.core.util.Context.NONE);
    }
}
