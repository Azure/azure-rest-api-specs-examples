import com.azure.core.util.Context;

/** Samples for StorageAccounts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountGetPropertiesCMKEnabled.json
     */
    /**
     * Sample code: StorageAccountGetPropertiesCMKEnabled.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountGetPropertiesCMKEnabled(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .getByResourceGroupWithResponse("res9407", "sto8596", null, Context.NONE);
    }
}
