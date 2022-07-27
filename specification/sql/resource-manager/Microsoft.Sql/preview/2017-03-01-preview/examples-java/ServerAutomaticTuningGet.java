import com.azure.core.util.Context;

/** Samples for ServerAutomaticTuning Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ServerAutomaticTuningGet.json
     */
    /**
     * Sample code: Get a server's automatic tuning settings.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAServerSAutomaticTuningSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerAutomaticTunings()
            .getWithResponse("default-sql-onebox", "testsvr11", Context.NONE);
    }
}
