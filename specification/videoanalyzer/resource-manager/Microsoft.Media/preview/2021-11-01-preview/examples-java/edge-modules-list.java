import com.azure.core.util.Context;

/** Samples for EdgeModules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/edge-modules-list.json
     */
    /**
     * Sample code: Lists the registered edge modules.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void listsTheRegisteredEdgeModules(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.edgeModules().list("testrg", "testaccount2", null, Context.NONE);
    }
}
