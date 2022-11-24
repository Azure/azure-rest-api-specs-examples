import com.azure.core.util.Context;

/** Samples for Accounts ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/ListKeys.json
     */
    /**
     * Sample code: List Keys.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listKeys(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listKeysWithResponse("myResourceGroup", "myAccount", Context.NONE);
    }
}
