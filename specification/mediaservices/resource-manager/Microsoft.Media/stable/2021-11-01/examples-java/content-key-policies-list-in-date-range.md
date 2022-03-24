Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ContentKeyPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/content-key-policies-list-in-date-range.json
     */
    /**
     * Sample code: Lists Content Key Policies with created and last modified filters.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsContentKeyPoliciesWithCreatedAndLastModifiedFilters(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .contentKeyPolicies()
            .list(
                "contoso",
                "contosomedia",
                "properties/lastModified gt 2016-06-01 and properties/created lt 2013-07-01",
                null,
                null,
                Context.NONE);
    }
}
```
