```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.videoanalyzer.models.CheckNameAvailabilityRequest;

/** Samples for Locations CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/accounts-check-name-availability.json
     */
    /**
     * Sample code: Check Name Availability.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .locations()
            .checkNameAvailabilityWithResponse(
                "japanwest",
                new CheckNameAvailabilityRequest().withName("contosotv").withType("videoAnalyzers"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.
