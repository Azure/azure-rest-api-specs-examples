```java
/** Samples for EdgeModules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/edge-modules-create.json
     */
    /**
     * Sample code: Registers an edge module.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void registersAnEdgeModule(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.edgeModules().define("edgeModule1").withExistingVideoAnalyzer("testrg", "testaccount2").create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.
