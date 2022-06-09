```java
import com.azure.core.util.Context;

/** Samples for DomainWhois Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/enrichment/GetWhoisByDomainName.json
     */
    /**
     * Sample code: Get whois information for a single domain name.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getWhoisInformationForASingleDomainName(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.domainWhois().getWithResponse("myRg", "microsoft.com", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
