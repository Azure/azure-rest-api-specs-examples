Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.AssetContainerPermission;
import com.azure.resourcemanager.mediaservices.models.ListContainerSasInput;
import java.time.OffsetDateTime;

/** Samples for Assets ListContainerSas. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-list-sas-urls.json
     */
    /**
     * Sample code: List Asset SAS URLs.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAssetSASURLs(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .assets()
            .listContainerSasWithResponse(
                "contoso",
                "contosomedia",
                "ClimbingMountBaker",
                new ListContainerSasInput()
                    .withPermissions(AssetContainerPermission.READ_WRITE)
                    .withExpiryTime(OffsetDateTime.parse("2018-01-01T10:00:00.007Z")),
                Context.NONE);
    }
}
```
