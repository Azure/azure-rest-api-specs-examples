import com.azure.core.util.Context;

/** Samples for FileShares Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesGet_Stats.json
     */
    /**
     * Sample code: GetShareStats.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getShareStats(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .getWithResponse("res9871", "sto6217", "share1634", "stats", null, Context.NONE);
    }
}
