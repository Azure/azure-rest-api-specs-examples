
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.BastionHostInner;
import com.azure.resourcemanager.network.models.BastionHostIpConfiguration;
import com.azure.resourcemanager.network.models.BastionSessionRecordingConfiguration;
import com.azure.resourcemanager.network.models.SessionRecordingIdentity;
import com.azure.resourcemanager.network.models.SessionRecordingIdentityType;
import java.util.Arrays;

/**
 * Samples for BastionHosts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/BastionHostPutWithSystemAssignedIdentityForSRConfig.json
     */
    /**
     * Sample code: Create or Update Bastion Host With System Assigned Identity for Session Recording Configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createOrUpdateBastionHostWithSystemAssignedIdentityForSessionRecordingConfiguration(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().createOrUpdate("rg1", "bastionhosttenant", new BastionHostInner()
            .withIpConfigurations(Arrays.asList(new BastionHostIpConfiguration().withName("bastionHostIpConfiguration")
                .withSubnet(new SubResource().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"))
                .withPublicIpAddress(new SubResource().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"))))
            .withEnableSessionRecording(true)
            .withSessionRecordingConfiguration(new BastionSessionRecordingConfiguration()
                .withIdentity(new SessionRecordingIdentity().withType(SessionRecordingIdentityType.SYSTEM_ASSIGNED))
                .withBlobContainerUri("https://mystorageaccount.blob.core.windows.net/mycontainer")),
            com.azure.core.util.Context.NONE);
    }
}
