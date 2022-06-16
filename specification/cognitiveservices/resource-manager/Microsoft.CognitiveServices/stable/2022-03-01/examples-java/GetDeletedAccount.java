import com.azure.core.util.Context;

/** Samples for DeletedAccounts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/GetDeletedAccount.json
     */
    /**
     * Sample code: Get Account.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getAccount(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deletedAccounts().getWithResponse("westus", "myResourceGroup", "myAccount", Context.NONE);
    }
}
