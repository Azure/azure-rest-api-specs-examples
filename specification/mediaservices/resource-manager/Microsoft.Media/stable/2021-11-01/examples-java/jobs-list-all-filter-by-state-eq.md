Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-list-all-filter-by-state-eq.json
     */
    /**
     * Sample code: Lists Jobs for the Transform filter by state equal.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsJobsForTheTransformFilterByStateEqual(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .jobs()
            .list(
                "contosoresources",
                "contosomedia",
                "exampleTransform",
                "properties/state eq Microsoft.Media.JobState'Processing'",
                null,
                Context.NONE);
    }
}
```
