```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.NamedValueContract;
import java.util.Arrays;

/** Samples for NamedValue Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateNamedValue.json
     */
    /**
     * Sample code: ApiManagementUpdateNamedValue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateNamedValue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        NamedValueContract resource =
            manager.namedValues().getWithResponse("rg1", "apimService1", "testprop2", Context.NONE).getValue();
        resource
            .update()
            .withTags(Arrays.asList("foo", "bar2"))
            .withDisplayName("prop3name")
            .withValue("propValue")
            .withSecret(false)
            .withIfMatch("*")
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
