Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-labservices_1.0.0-beta.2/sdk/labservices/azure-resourcemanager-labservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.labservices.models.EnableState;
import com.azure.resourcemanager.labservices.models.Image;

/** Samples for Images Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Images/patchImage.json
     */
    /**
     * Sample code: patchImage.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void patchImage(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        Image resource =
            manager.images().getWithResponse("testrg123", "testlabplan", "image1", Context.NONE).getValue();
        resource.update().withEnabledState(EnableState.ENABLED).apply();
    }
}
```
