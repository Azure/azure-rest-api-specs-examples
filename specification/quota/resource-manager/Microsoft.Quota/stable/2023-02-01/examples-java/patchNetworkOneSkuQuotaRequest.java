import com.azure.resourcemanager.quota.models.CurrentQuotaLimitBase;
import com.azure.resourcemanager.quota.models.LimitObject;
import com.azure.resourcemanager.quota.models.QuotaProperties;
import com.azure.resourcemanager.quota.models.ResourceName;

/** Samples for Quota Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/patchNetworkOneSkuQuotaRequest.json
     */
    /**
     * Sample code: Quotas_Request_PatchForNetwork.
     *
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasRequestPatchForNetwork(com.azure.resourcemanager.quota.QuotaManager manager) {
        CurrentQuotaLimitBase resource =
            manager
                .quotas()
                .getWithResponse(
                    "MinPublicIpInterNetworkPrefixLength",
                    "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Network/locations/eastus",
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(
                new QuotaProperties()
                    .withLimit(new LimitObject().withValue(10))
                    .withName(new ResourceName().withValue("MinPublicIpInterNetworkPrefixLength"))
                    .withResourceType("MinPublicIpInterNetworkPrefixLength"))
            .apply();
    }
}
