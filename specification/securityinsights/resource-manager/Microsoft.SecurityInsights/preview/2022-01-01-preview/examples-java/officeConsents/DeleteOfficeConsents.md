Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for OfficeConsents Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/officeConsents/DeleteOfficeConsents.json
     */
    /**
     * Sample code: Delete an office consent.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAnOfficeConsent(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .officeConsents()
            .deleteWithResponse("myRg", "myWorkspace", "04e5fd05-ff86-4b97-b8d2-1c20933cb46c", Context.NONE);
    }
}
```
