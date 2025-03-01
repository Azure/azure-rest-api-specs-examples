
/**
 * Samples for FileShares Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileSharesGet_ProvisionedV2.
     * json
     */
    /**
     * Sample code: GetShareProvisionedV2.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getShareProvisionedV2(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileShares().getWithResponse("res9871", "sto6217",
            "share1634", null, null, com.azure.core.util.Context.NONE);
    }
}
