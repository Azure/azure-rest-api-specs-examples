
/**
 * Samples for DeletedAccounts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/DeletedAccountGet.json
     */
    /**
     * Sample code: DeletedAccountGet.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deletedAccountGet(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getDeletedAccounts().getWithResponse("sto1125", "eastus",
            com.azure.core.util.Context.NONE);
    }
}
