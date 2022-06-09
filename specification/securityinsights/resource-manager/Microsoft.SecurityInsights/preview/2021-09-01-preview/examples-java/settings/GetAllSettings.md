```java
import com.azure.core.util.Context;

/** Samples for ProductSettings List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/settings/GetAllSettings.json
     */
    /**
     * Sample code: Get all settings.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllSettings(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.productSettings().listWithResponse("myRg", "myWorkspace", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
