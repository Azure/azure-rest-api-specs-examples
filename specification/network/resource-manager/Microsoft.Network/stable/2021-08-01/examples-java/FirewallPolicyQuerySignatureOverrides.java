import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.FilterItems;
import com.azure.resourcemanager.network.models.FirewallPolicyIdpsQuerySortOrder;
import com.azure.resourcemanager.network.models.IdpsQueryObject;
import com.azure.resourcemanager.network.models.OrderBy;
import java.util.Arrays;

/** Samples for FirewallPolicyIdpsSignatures List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/FirewallPolicyQuerySignatureOverrides.json
     */
    /**
     * Sample code: query signature overrides.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void querySignatureOverrides(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFirewallPolicyIdpsSignatures()
            .listWithResponse(
                "rg1",
                "firewallPolicy",
                new IdpsQueryObject()
                    .withFilters(Arrays.asList(new FilterItems().withField("Mode").withValues(Arrays.asList("Deny"))))
                    .withSearch("")
                    .withOrderBy(
                        new OrderBy().withField("severity").withOrder(FirewallPolicyIdpsQuerySortOrder.ASCENDING))
                    .withResultsPerPage(20)
                    .withSkip(0),
                Context.NONE);
    }
}
