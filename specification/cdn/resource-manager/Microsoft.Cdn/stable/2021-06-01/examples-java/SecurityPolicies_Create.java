import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.fluent.models.SecurityPolicyInner;
import com.azure.resourcemanager.cdn.models.ActivatedResourceReference;
import com.azure.resourcemanager.cdn.models.ResourceReference;
import com.azure.resourcemanager.cdn.models.SecurityPolicyWebApplicationFirewallAssociation;
import com.azure.resourcemanager.cdn.models.SecurityPolicyWebApplicationFirewallParameters;
import java.util.Arrays;

/** Samples for SecurityPolicies Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/SecurityPolicies_Create.json
     */
    /**
     * Sample code: SecurityPolicies_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void securityPoliciesCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getSecurityPolicies()
            .create(
                "RG",
                "profile1",
                "securityPolicy1",
                new SecurityPolicyInner()
                    .withParameters(
                        new SecurityPolicyWebApplicationFirewallParameters()
                            .withWafPolicy(
                                new ResourceReference()
                                    .withId(
                                        "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoorwebapplicationfirewallpolicies/wafTest"))
                            .withAssociations(
                                Arrays
                                    .asList(
                                        new SecurityPolicyWebApplicationFirewallAssociation()
                                            .withDomains(
                                                Arrays
                                                    .asList(
                                                        new ActivatedResourceReference()
                                                            .withId(
                                                                "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/afddomains/testdomain1"),
                                                        new ActivatedResourceReference()
                                                            .withId(
                                                                "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/afddomains/testdomain2")))
                                            .withPatternsToMatch(Arrays.asList("/*"))))),
                Context.NONE);
    }
}
