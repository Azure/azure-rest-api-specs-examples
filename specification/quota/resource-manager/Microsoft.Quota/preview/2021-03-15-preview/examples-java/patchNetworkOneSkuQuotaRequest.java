import com.azure.core.util.Context;
import com.azure.resourcemanager.quota.models.CurrentQuotaLimitBase;
import com.azure.resourcemanager.quota.models.LimitObject;
import com.azure.resourcemanager.quota.models.QuotaProperties;
import com.azure.resourcemanager.quota.models.ResourceName;

/** Samples for Quota Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/patchNetworkOneSkuQuotaRequest.json
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
                    Context.NONE)
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
