
/**
 * Samples for Accounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * GetAccount.json
     */
    /**
     * Sample code: Get Account.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getAccount(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().getByResourceGroupWithResponse("myResourceGroup", "myAccount",
            com.azure.core.util.Context.NONE);
    }
}
