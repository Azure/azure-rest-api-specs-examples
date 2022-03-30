Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import java.util.Arrays;

/** Samples for NamedValue CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateNamedValue.json
     */
    /**
     * Sample code: ApiManagementCreateNamedValue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateNamedValue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .namedValues()
            .define("testprop2")
            .withExistingService("rg1", "apimService1")
            .withTags(Arrays.asList("foo", "bar"))
            .withDisplayName("prop3name")
            .withValue("propValue")
            .withSecret(false)
            .create();
    }
}
```
