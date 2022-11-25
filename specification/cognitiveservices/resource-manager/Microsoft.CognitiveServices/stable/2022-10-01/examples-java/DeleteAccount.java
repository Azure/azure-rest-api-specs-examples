import com.azure.core.util.Context;

/** Samples for Accounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/DeleteAccount.json
     */
    /**
     * Sample code: Delete Account.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteAccount(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().delete("myResourceGroup", "PropTest01", Context.NONE);
    }
}
