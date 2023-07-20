/** Samples for Accounts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/ListAccountsByResourceGroup.json
     */
    /**
     * Sample code: List Accounts by Resource Group.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listAccountsByResourceGroup(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
