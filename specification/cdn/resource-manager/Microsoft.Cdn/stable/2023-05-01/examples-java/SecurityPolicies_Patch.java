
import com.azure.resourcemanager.cdn.models.ActivatedResourceReference;
import com.azure.resourcemanager.cdn.models.ResourceReference;
import com.azure.resourcemanager.cdn.models.SecurityPolicyUpdateParameters;
import com.azure.resourcemanager.cdn.models.SecurityPolicyWebApplicationFirewallAssociation;
import com.azure.resourcemanager.cdn.models.SecurityPolicyWebApplicationFirewallParameters;
import java.util.Arrays;

/** Samples for SecurityPolicies Patch. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/SecurityPolicies_Patch.json
     */
    /**
     * Sample code: SecurityPolicies_Patch.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void securityPoliciesPatch(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getSecurityPolicies().patch("RG", "profile1", "securityPolicy1",
            new SecurityPolicyUpdateParameters().withParameters(new SecurityPolicyWebApplicationFirewallParameters()
                .withWafPolicy(new ResourceReference().withId(
                    "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoorwebapplicationfirewallpolicies/wafTest"))
                .withAssociations(Arrays.asList(new SecurityPolicyWebApplicationFirewallAssociation()
                    .withDomains(Arrays.asList(new ActivatedResourceReference().withId(
                        "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain1"),
                        new ActivatedResourceReference().withId(
                            "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain2")))
                    .withPatternsToMatch(Arrays.asList("/*"))))),
            com.azure.core.util.Context.NONE);
    }
}
