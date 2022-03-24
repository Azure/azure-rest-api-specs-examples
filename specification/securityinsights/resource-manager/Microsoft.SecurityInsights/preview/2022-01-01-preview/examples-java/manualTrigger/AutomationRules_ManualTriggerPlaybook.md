Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.2/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Incidents RunPlaybook. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/manualTrigger/AutomationRules_ManualTriggerPlaybook.json
     */
    /**
     * Sample code: AutomationRules_ManualTriggerPlaybook.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void automationRulesManualTriggerPlaybook(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidents()
            .runPlaybookWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ar4", null, Context.NONE);
    }
}
```
