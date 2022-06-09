```java
import com.azure.core.util.Context;

/** Samples for EdgeModules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/edge-modules-delete.json
     */
    /**
     * Sample code: Deletes an edge module registration.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void deletesAnEdgeModuleRegistration(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.edgeModules().deleteWithResponse("testrg", "testaccount2", "edgeModule1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.
