
import com.azure.resourcemanager.securityinsights.models.EnrichmentDomainBody;
import com.azure.resourcemanager.securityinsights.models.EnrichmentType;

/**
 * Samples for ResourceProvider ListWhoisByDomain.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/enrichment/GetWhoisWithWorkspaceByDomainName.json
     */
    /**
     * Sample code: Get whois information for a single domain name.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getWhoisInformationForASingleDomainName(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.resourceProviders().listWhoisByDomainWithResponse("myRg", "myWorkspace", EnrichmentType.MAIN,
            new EnrichmentDomainBody().withDomain("microsoft.com"), com.azure.core.util.Context.NONE);
    }
}
