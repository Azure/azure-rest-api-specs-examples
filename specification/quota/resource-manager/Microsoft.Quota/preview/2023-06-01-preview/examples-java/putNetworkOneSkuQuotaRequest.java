
import com.azure.resourcemanager.quota.models.LimitObject;
import com.azure.resourcemanager.quota.models.QuotaProperties;
import com.azure.resourcemanager.quota.models.ResourceName;

/**
 * Samples for Quota CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/
     * putNetworkOneSkuQuotaRequest.json
     */
    /**
     * Sample code: Quotas_PutRequest_ForNetwork.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasPutRequestForNetwork(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.quotas().define("MinPublicIpInterNetworkPrefixLength")
            .withExistingScope(
                "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Network/locations/eastus")
            .withProperties(new QuotaProperties().withLimit(new LimitObject().withValue(10))
                .withName(new ResourceName().withValue("MinPublicIpInterNetworkPrefixLength"))
                .withResourceType("MinPublicIpInterNetworkPrefixLength"))
            .create();
    }
}
