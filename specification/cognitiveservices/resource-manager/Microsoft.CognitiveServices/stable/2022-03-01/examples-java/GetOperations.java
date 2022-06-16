import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/GetOperations.json
     */
    /**
     * Sample code: Get Operations.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getOperations(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.operations().list(Context.NONE);
    }
}
