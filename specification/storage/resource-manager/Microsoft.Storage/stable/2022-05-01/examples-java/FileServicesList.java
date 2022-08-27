import com.azure.core.util.Context;

/** Samples for FileServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/FileServicesList.json
     */
    /**
     * Sample code: ListFileServices.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listFileServices(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileServices()
            .listWithResponse("res9290", "sto1590", Context.NONE);
    }
}
