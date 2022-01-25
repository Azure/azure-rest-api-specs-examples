Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AutomationRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/automationRules/GetAllAutomationRules.json
     */
    /**
     * Sample code: Get all automation rules.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllAutomationRules(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.automationRules().list("myRg", "myWorkspace", Context.NONE);
    }
}
```
