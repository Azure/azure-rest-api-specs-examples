
import com.azure.resourcemanager.enclave.models.CheckAddressSpaceAvailabilityRequest;
import com.azure.resourcemanager.enclave.models.EnclaveVirtualNetworkModel;
import com.azure.resourcemanager.enclave.models.SubnetConfiguration;
import java.util.Arrays;

/**
 * Samples for Community CheckAddressSpaceAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Community_PostCheckAddressSpaceAvailability.json
     */
    /**
     * Sample code: Community_CheckAddressSpaceAvailability.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        communityCheckAddressSpaceAvailability(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.communities().checkAddressSpaceAvailabilityWithResponse("rgopenapi", "TestMyCommunity",
            new CheckAddressSpaceAvailabilityRequest().withCommunityResourceId(
                "/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/Microsoft.Mission/communities/TestMyCommunity")
                .withEnclaveVirtualNetwork(
                    new EnclaveVirtualNetworkModel().withNetworkSize("small").withCustomCidrRange("10.0.0.0/24")
                        .withSubnetConfigurations(
                            Arrays.asList(new SubnetConfiguration().withSubnetName("test").withNetworkPrefixSize(26)))
                        .withAllowSubnetCommunication(true)),
            com.azure.core.util.Context.NONE);
    }
}
