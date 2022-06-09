```java
/** Samples for Videos CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-create.json
     */
    /**
     * Sample code: Register video entity.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void registerVideoEntity(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .videos()
            .define("video1")
            .withExistingVideoAnalyzer("testrg", "testaccount2")
            .withTitle("Sample Title 1")
            .withDescription("Sample Description 1")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.
