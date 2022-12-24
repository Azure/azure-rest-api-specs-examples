import com.azure.core.util.Context;

/** Samples for Table Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/TableOperationPut.json
     */
    /**
     * Sample code: TableOperationPut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tableOperationPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getTables()
            .createWithResponse("res3376", "sto328", "table6185", null, Context.NONE);
    }
}
