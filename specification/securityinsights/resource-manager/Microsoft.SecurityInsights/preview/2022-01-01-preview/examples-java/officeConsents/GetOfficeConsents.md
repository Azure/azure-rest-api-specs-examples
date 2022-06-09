```java
import com.azure.core.util.Context;

/** Samples for OfficeConsents List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/officeConsents/GetOfficeConsents.json
     */
    /**
     * Sample code: Get all office consents.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllOfficeConsents(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.officeConsents().list("myRg", "myWorkspace", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
