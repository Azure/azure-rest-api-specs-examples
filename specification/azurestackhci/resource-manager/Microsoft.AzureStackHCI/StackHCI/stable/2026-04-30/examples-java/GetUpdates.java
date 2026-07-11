
/**
 * Samples for Updates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/GetUpdates.json
     */
    /**
     * Sample code: Get a specific update.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getASpecificUpdate(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updates().getWithResponse("testrg", "testcluster", "Microsoft4.2203.2.32",
            com.azure.core.util.Context.NONE);
    }
}
