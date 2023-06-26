import com.azure.core.util.Context;

/** Samples for Storages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Storages_List.json
     */
    /**
     * Sample code: Storages_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storagesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getStorages()
            .list("myResourceGroup", "myservice", Context.NONE);
    }
}
