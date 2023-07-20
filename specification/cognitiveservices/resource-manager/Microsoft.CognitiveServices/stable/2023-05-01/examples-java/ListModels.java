/** Samples for Models List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/ListModels.json
     */
    /**
     * Sample code: ListModels.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listModels(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.models().list("WestUS", com.azure.core.util.Context.NONE);
    }
}
