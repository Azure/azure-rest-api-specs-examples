import com.azure.core.util.Context;

/** Samples for Accounts ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/ListSkus.json
     */
    /**
     * Sample code: List SKUs.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listSKUs(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listSkusWithResponse("myResourceGroup", "myAccount", Context.NONE);
    }
}
