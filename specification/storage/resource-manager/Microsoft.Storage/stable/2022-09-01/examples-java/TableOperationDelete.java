import com.azure.core.util.Context;

/** Samples for Table Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/TableOperationDelete.json
     */
    /**
     * Sample code: TableOperationDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tableOperationDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getTables()
            .deleteWithResponse("res3376", "sto328", "table6185", Context.NONE);
    }
}
