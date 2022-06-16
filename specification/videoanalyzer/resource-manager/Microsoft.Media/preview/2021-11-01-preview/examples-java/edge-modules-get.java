import com.azure.core.util.Context;

/** Samples for EdgeModules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/edge-modules-get.json
     */
    /**
     * Sample code: Gets edge module registration.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getsEdgeModuleRegistration(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.edgeModules().getWithResponse("testrg", "testaccount2", "edgeModule1", Context.NONE);
    }
}
