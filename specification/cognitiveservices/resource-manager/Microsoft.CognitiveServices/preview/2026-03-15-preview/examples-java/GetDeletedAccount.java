
/**
 * Samples for DeletedAccounts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/GetDeletedAccount.json
     */
    /**
     * Sample code: Get Account.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getAccount(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deletedAccounts().getWithResponse("westus", "myResourceGroup", "myAccount",
            com.azure.core.util.Context.NONE);
    }
}
