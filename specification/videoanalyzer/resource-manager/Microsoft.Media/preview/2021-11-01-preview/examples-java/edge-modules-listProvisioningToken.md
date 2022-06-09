```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.videoanalyzer.models.ListProvisioningTokenInput;
import java.time.OffsetDateTime;

/** Samples for EdgeModules ListProvisioningToken. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/edge-modules-listProvisioningToken.json
     */
    /**
     * Sample code: Generate the Provisioning token for an edge module registration.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void generateTheProvisioningTokenForAnEdgeModuleRegistration(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .edgeModules()
            .listProvisioningTokenWithResponse(
                "testrg",
                "testaccount2",
                "edgeModule1",
                new ListProvisioningTokenInput()
                    .withExpirationDate(OffsetDateTime.parse("2023-01-23T11:04:49.0526841-08:00")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.
