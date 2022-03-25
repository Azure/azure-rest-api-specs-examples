Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.2/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/dataConnectors/GetAmazonWebServicesCloudTrailById.json
     */
    /**
     * Sample code: Get an AwsCloudTrail data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnAwsCloudTrailDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "c345bf40-8509-4ed2-b947-50cb773aaf04", Context.NONE);
    }
}
```
