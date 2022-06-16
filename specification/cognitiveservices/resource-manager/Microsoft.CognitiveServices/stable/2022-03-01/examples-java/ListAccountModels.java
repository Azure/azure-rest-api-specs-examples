import com.azure.core.util.Context;

/** Samples for Accounts ListModels. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/ListAccountModels.json
     */
    /**
     * Sample code: List AccountModels.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listAccountModels(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listModels("resourceGroupName", "accountName", Context.NONE);
    }
}
