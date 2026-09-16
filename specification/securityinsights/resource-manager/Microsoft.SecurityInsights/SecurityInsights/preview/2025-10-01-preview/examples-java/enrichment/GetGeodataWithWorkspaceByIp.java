
import com.azure.resourcemanager.securityinsights.models.EnrichmentIpAddressBody;
import com.azure.resourcemanager.securityinsights.models.EnrichmentType;

/**
 * Samples for ResourceProvider ListGeodataByIp.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/enrichment/GetGeodataWithWorkspaceByIp.json
     */
    /**
     * Sample code: Get geodata for a single IP address.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getGeodataForASingleIPAddress(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.resourceProviders().listGeodataByIpWithResponse("myRg", "myWorkspace", EnrichmentType.MAIN,
            new EnrichmentIpAddressBody().withIpAddress("1.2.3.4"), com.azure.core.util.Context.NONE);
    }
}
