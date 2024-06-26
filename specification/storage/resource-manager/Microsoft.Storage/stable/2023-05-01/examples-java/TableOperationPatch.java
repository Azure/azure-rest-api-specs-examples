
/**
 * Samples for Table Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/TableOperationPatch.json
     */
    /**
     * Sample code: TableOperationPatch.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tableOperationPatch(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getTables().updateWithResponse("res3376", "sto328",
            "table6185", null, com.azure.core.util.Context.NONE);
    }
}
