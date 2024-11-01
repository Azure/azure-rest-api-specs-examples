
/**
 * Samples for AccountFilters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/accountFilters-
     * list-all.json
     */
    /**
     * Sample code: List all Account Filters.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllAccountFilters(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.accountFilters().list("contosorg", "contosomedia", com.azure.core.util.Context.NONE);
    }
}
