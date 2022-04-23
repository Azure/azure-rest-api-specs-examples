Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-servicelinker_1.0.0-beta.1/sdk/servicelinker/azure-resourcemanager-servicelinker/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Linker Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/ValidateLinkSuccess.json
     */
    /**
     * Sample code: ValidateLinkSuccess.
     *
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void validateLinkSuccess(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager
            .linkers()
            .validate(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
                "linkName",
                Context.NONE);
    }
}
```
