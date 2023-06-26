import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.StorageResourceInner;
import com.azure.resourcemanager.appplatform.models.StorageAccount;

/** Samples for Storages CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Storages_CreateOrUpdate.json
     */
    /**
     * Sample code: Storages_CreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storagesCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getStorages()
            .createOrUpdate(
                "myResourceGroup",
                "myservice",
                "mystorage",
                new StorageResourceInner()
                    .withProperties(
                        new StorageAccount()
                            .withAccountName("storage-account-name")
                            .withAccountKey("account-key-of-storage-account")),
                Context.NONE);
    }
}
