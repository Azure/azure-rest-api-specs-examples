
/**
 * Samples for Storages Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Storages_Delete.json
     */
    /**
     * Sample code: Storages_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storagesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getStorages().delete("myResourceGroup", "myservice",
            "mystorage", com.azure.core.util.Context.NONE);
    }
}
