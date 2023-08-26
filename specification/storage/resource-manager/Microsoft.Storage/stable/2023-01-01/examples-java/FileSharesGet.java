/** Samples for FileShares Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/FileSharesGet.json
     */
    /**
     * Sample code: GetShares.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getShares(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .getWithResponse("res9871", "sto6217", "share1634", null, null, com.azure.core.util.Context.NONE);
    }
}
