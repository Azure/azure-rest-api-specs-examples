Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-labservices_1.0.0-beta.2/sdk/labservices/azure-resourcemanager-labservices/README.md) on how to add the SDK to your project and authenticate.

```java
import java.time.Duration;

/** Samples for Users CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Users/putUser.json
     */
    /**
     * Sample code: putUser.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void putUser(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager
            .users()
            .define("testuser")
            .withExistingLab("testrg123", "testlab")
            .withEmail("testuser@contoso.com")
            .withAdditionalUsageQuota(Duration.parse("20:00"))
            .create();
    }
}
```
