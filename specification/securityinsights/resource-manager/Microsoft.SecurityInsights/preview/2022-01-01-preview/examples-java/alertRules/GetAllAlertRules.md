```java
import com.azure.core.util.Context;

/** Samples for AlertRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/alertRules/GetAllAlertRules.json
     */
    /**
     * Sample code: Get all alert rules.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllAlertRules(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.alertRules().list("myRg", "myWorkspace", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
