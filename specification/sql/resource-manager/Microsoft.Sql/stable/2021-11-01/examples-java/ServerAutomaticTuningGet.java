
/**
 * Samples for ServerAutomaticTuning Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerAutomaticTuningGet.json
     */
    /**
     * Sample code: Get a server's automatic tuning settings.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAServerSAutomaticTuningSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAutomaticTunings().getWithResponse("default-sql-onebox",
            "testsvr11", com.azure.core.util.Context.NONE);
    }
}
