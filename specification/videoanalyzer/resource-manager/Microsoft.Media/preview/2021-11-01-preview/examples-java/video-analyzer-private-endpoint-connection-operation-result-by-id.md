```java
import com.azure.core.util.Context;

/** Samples for OperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-private-endpoint-connection-operation-result-by-id.json
     */
    /**
     * Sample code: Get status of private endpoint connection asynchronous operation.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getStatusOfPrivateEndpointConnectionAsynchronousOperation(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .operationResults()
            .getWithResponse(
                "contoso",
                "contososports",
                "6FBA62C4-99B5-4FF8-9826-FC4744A8864F",
                "10000000-0000-0000-0000-000000000000",
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.
