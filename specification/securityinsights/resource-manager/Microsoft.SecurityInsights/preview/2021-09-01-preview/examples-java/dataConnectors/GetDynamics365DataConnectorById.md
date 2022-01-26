Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/GetDynamics365DataConnectorById.json
     */
    /**
     * Sample code: Get a Dynamics365 data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getADynamics365DataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "c2541efb-c9a6-47fe-9501-87d1017d1512", Context.NONE);
    }
}
```
