
/** Samples for TableServices List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/TableServicesList.json
     */
    /**
     * Sample code: TableServicesList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tableServicesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getTableServices().listWithResponse("res9290", "sto1590",
            com.azure.core.util.Context.NONE);
    }
}
