Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for IpGeodata Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/enrichment/GetGeodataByIp.json
     */
    /**
     * Sample code: Get geodata for a single IP address.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getGeodataForASingleIPAddress(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.ipGeodatas().getWithResponse("myRg", "1.2.3.4", Context.NONE);
    }
}
```
