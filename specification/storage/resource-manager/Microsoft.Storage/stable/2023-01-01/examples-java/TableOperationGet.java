
/** Samples for Table Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/TableOperationGet.json
     */
    /**
     * Sample code: TableOperationGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tableOperationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getTables().getWithResponse("res3376", "sto328", "table6185",
            com.azure.core.util.Context.NONE);
    }
}
