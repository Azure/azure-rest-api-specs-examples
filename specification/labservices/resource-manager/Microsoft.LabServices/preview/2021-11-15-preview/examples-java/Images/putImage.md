```java
import com.azure.resourcemanager.labservices.models.EnableState;

/** Samples for Images CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Images/putImage.json
     */
    /**
     * Sample code: putImage.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void putImage(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager
            .images()
            .define("image1")
            .withExistingLabPlan("testrg123", "testlabplan")
            .withEnabledState(EnableState.ENABLED)
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-labservices_1.0.0-beta.2/sdk/labservices/azure-resourcemanager-labservices/README.md) on how to add the SDK to your project and authenticate.
