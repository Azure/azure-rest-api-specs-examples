import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.VpnServerConfigurationPolicyGroupInner;
import com.azure.resourcemanager.network.models.VpnPolicyMemberAttributeType;
import com.azure.resourcemanager.network.models.VpnServerConfigurationPolicyGroupMember;
import java.util.Arrays;

/** Samples for ConfigurationPolicyGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ConfigurationPolicyGroupPut.json
     */
    /**
     * Sample code: ConfigurationPolicyGroupPut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configurationPolicyGroupPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getConfigurationPolicyGroups()
            .createOrUpdate(
                "rg1",
                "vpnServerConfiguration1",
                "policyGroup1",
                new VpnServerConfigurationPolicyGroupInner()
                    .withIsDefault(true)
                    .withPriority(0)
                    .withPolicyMembers(
                        Arrays
                            .asList(
                                new VpnServerConfigurationPolicyGroupMember()
                                    .withName("policy1")
                                    .withAttributeType(VpnPolicyMemberAttributeType.RADIUS_AZURE_GROUP_ID)
                                    .withAttributeValue("6ad1bd08"),
                                new VpnServerConfigurationPolicyGroupMember()
                                    .withName("policy2")
                                    .withAttributeType(VpnPolicyMemberAttributeType.CERTIFICATE_GROUP_ID)
                                    .withAttributeValue("red.com"))),
                Context.NONE);
    }
}
