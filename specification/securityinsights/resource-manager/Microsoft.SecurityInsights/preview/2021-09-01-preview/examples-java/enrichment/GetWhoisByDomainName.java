import com.azure.core.util.Context;

/** Samples for DomainWhois Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/enrichment/GetWhoisByDomainName.json
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
