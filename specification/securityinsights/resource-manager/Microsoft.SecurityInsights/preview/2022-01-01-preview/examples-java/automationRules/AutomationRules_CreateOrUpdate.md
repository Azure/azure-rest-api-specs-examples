```java
import com.azure.resourcemanager.securityinsights.models.AutomationRuleAction;
import com.azure.resourcemanager.securityinsights.models.AutomationRuleTriggeringLogic;
import java.util.List;

/** Samples for AutomationRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/automationRules/AutomationRules_CreateOrUpdate.json
     */
    /**
     * Sample code: AutomationRules_CreateOrUpdate.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void automationRulesCreateOrUpdate(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .automationRules()
            .define("73e01a99-5cd7-4139-a149-9f2736ff2ab5")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withDisplayName((String) null)
            .withOrder(0)
            .withTriggeringLogic((AutomationRuleTriggeringLogic) null)
            .withActions((List<AutomationRuleAction>) null)
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
