Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/GetMicrosoftDefenderAdvancedThreatProtectionById.json
     */
    /**
     * Sample code: Get a MDATP data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAMDATPDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "06b3ccb8-1384-4bcc-aec7-852f6d57161b", Context.NONE);
    }
}
```
