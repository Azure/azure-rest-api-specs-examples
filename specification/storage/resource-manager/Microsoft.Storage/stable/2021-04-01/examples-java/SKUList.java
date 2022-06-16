import com.azure.core.util.Context;

/** Samples for Skus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/SKUList.json
     */
    /**
     * Sample code: SkuList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void skuList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getSkus().list(Context.NONE);
    }
}
