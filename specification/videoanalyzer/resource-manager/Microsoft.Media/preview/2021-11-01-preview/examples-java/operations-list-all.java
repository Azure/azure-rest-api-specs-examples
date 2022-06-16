import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/operations-list-all.json
     */
    /**
     * Sample code: List Operations.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void listOperations(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.operations().listWithResponse(Context.NONE);
    }
}
