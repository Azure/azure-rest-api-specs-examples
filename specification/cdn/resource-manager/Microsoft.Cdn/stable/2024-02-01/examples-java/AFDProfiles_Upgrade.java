
import com.azure.resourcemanager.cdn.models.ProfileChangeSkuWafMapping;
import com.azure.resourcemanager.cdn.models.ProfileUpgradeParameters;
import com.azure.resourcemanager.cdn.models.ResourceReference;
import java.util.Arrays;

/**
 * Samples for AfdProfiles Upgrade.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDProfiles_Upgrade.json
     */
    /**
     * Sample code: AFDProfiles_Upgrade.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDProfilesUpgrade(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdProfiles().upgrade("RG", "profile1",
            new ProfileUpgradeParameters().withWafMappingList(Arrays.asList(new ProfileChangeSkuWafMapping()
                .withSecurityPolicyName("securityPolicy1")
                .withChangeToWafPolicy(new ResourceReference().withId(
                    "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoorwebapplicationfirewallpolicies/waf2")))),
            com.azure.core.util.Context.NONE);
    }
}
