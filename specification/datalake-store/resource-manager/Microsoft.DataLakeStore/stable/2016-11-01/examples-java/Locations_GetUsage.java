/** Samples for Locations GetUsage. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Locations_GetUsage.json
     */
    /**
     * Sample code: UsageList.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void usageList(com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager.locations().getUsage("WestUS", com.azure.core.util.Context.NONE);
    }
}
