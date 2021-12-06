Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.GalleryImageVersionInner;
import com.azure.resourcemanager.compute.models.GalleryArtifactVersionSource;
import com.azure.resourcemanager.compute.models.GalleryImageVersionPublishingProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
import com.azure.resourcemanager.compute.models.ReplicationMode;
import com.azure.resourcemanager.compute.models.TargetRegion;
import java.util.Arrays;

/** Samples for GalleryImageVersions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/CreateOrUpdateASimpleGalleryImageVersionWithShallowReplicationMode.json
     */
    /**
     * Sample code: Create or update a simple Gallery Image Version using shallow replication mode.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateASimpleGalleryImageVersionUsingShallowReplicationMode(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryImageVersions()
            .createOrUpdate(
                "myResourceGroup",
                "myGalleryName",
                "myGalleryImageName",
                "1.0.0",
                new GalleryImageVersionInner()
                    .withLocation("West US")
                    .withPublishingProfile(
                        new GalleryImageVersionPublishingProfile()
                            .withTargetRegions(
                                Arrays.asList(new TargetRegion().withName("West US").withRegionalReplicaCount(1)))
                            .withReplicationMode(ReplicationMode.SHALLOW))
                    .withStorageProfile(
                        new GalleryImageVersionStorageProfile()
                            .withSource(
                                new GalleryArtifactVersionSource()
                                    .withId(
                                        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"))),
                Context.NONE);
    }
}
```
