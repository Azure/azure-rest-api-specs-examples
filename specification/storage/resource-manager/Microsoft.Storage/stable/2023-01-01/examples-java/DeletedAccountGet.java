
/** Samples for DeletedAccounts Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/DeletedAccountGet.json
     */
    /**
     * Sample code: DeletedAccountGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletedAccountGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getDeletedAccounts().getWithResponse("sto1125", "eastus",
            com.azure.core.util.Context.NONE);
    }
}
