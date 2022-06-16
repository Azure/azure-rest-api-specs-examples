import com.azure.core.util.Context;

/** Samples for Table List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/TableOperationList.json
     */
    /**
     * Sample code: TableOperationList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tableOperationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getTables().list("res9290", "sto328", Context.NONE);
    }
}
