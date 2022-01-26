Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Entities Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entities/GetRegistryValueEntityById.json
     */
    /**
     * Sample code: Get a registry value entity.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getARegistryValueEntity(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "dc44bd11-b348-4d76-ad29-37bf7aa41356", Context.NONE);
    }
}
```
