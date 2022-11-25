import com.azure.core.util.Context;

/** Samples for ResourceSkus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/GetSkus.json
     */
    /**
     * Sample code: Regenerate Keys.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void regenerateKeys(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.resourceSkus().list(Context.NONE);
    }
}
