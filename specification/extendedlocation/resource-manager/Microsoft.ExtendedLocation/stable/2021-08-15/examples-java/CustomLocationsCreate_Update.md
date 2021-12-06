Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-extendedlocation_1.0.0-beta.1/sdk/extendedlocation/azure-resourcemanager-extendedlocation/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.extendedlocation.models.CustomLocationPropertiesAuthentication;
import com.azure.resourcemanager.extendedlocation.models.Identity;
import com.azure.resourcemanager.extendedlocation.models.ResourceIdentityType;
import java.util.Arrays;

/** Samples for CustomLocations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsCreate_Update.json
     */
    /**
     * Sample code: Create/Update Custom Location.
     *
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void createUpdateCustomLocation(
        com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager
            .customLocations()
            .define("customLocation01")
            .withRegion("West US")
            .withExistingResourceGroup("testresourcegroup")
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withAuthentication(
                new CustomLocationPropertiesAuthentication().withType("KubeConfig").withValue("<base64 KubeConfig>"))
            .withClusterExtensionIds(
                Arrays
                    .asList(
                        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedCluster/someCluster/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension"))
            .withDisplayName("customLocationLocation01")
            .withHostResourceId(
                "/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01")
            .withNamespace("namespace01")
            .create();
    }
}
```
