```java
import com.azure.core.util.Context;

/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-list-all.json
     */
    /**
     * Sample code: Lists all of the Jobs for the Transform.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsAllOfTheJobsForTheTransform(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.jobs().list("contosoresources", "contosomedia", "exampleTransform", null, null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.
