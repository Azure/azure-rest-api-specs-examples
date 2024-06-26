import com.azure.resourcemanager.hybridnetwork.models.IpAllocationMethod;
import com.azure.resourcemanager.hybridnetwork.models.IpVersion;
import com.azure.resourcemanager.hybridnetwork.models.LinuxConfiguration;
import com.azure.resourcemanager.hybridnetwork.models.NetworkFunctionVendorConfiguration;
import com.azure.resourcemanager.hybridnetwork.models.NetworkInterface;
import com.azure.resourcemanager.hybridnetwork.models.NetworkInterfaceIpConfiguration;
import com.azure.resourcemanager.hybridnetwork.models.OsProfile;
import com.azure.resourcemanager.hybridnetwork.models.SshConfiguration;
import com.azure.resourcemanager.hybridnetwork.models.SshPublicKey;
import com.azure.resourcemanager.hybridnetwork.models.VMSwitchType;
import com.azure.resourcemanager.hybridnetwork.models.VendorProvisioningState;
import java.util.Arrays;

/** Samples for VendorNetworkFunctions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorNfCreate.json
     */
    /**
     * Sample code: Create or update vendor network function sub resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createOrUpdateVendorNetworkFunctionSubResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager
            .vendorNetworkFunctions()
            .define("testServiceKey")
            .withExistingVendor("eastus", "testVendor")
            .withVendorProvisioningState(VendorProvisioningState.PROVISIONING)
            .withNetworkFunctionVendorConfigurations(
                Arrays
                    .asList(
                        new NetworkFunctionVendorConfiguration()
                            .withRoleName("testRole")
                            .withOsProfile(
                                new OsProfile()
                                    .withAdminUsername("dummyuser")
                                    .withLinuxConfiguration(
                                        new LinuxConfiguration()
                                            .withSsh(
                                                new SshConfiguration()
                                                    .withPublicKeys(
                                                        Arrays
                                                            .asList(
                                                                new SshPublicKey()
                                                                    .withPath("home/user/.ssh/authorized_keys")
                                                                    .withKeyData(
                                                                        "ssh-rsa"
                                                                            + " AAAAB3NzaC1yc2EAAAABIwAAAgEAwrr66r8n6B8Y0zMF3dOpXEapIQD9DiYQ6D6/zwor9o39jSkHNiMMER/GETBbzP83LOcekm02aRjo55ArO7gPPVvCXbrirJu9pkm4AC4BBre5xSLS="
                                                                            + " user@constoso-DSH")))))
                                    .withCustomData("base-64 encoded string of custom data"))
                            .withNetworkInterfaces(
                                Arrays
                                    .asList(
                                        new NetworkInterface()
                                            .withNetworkInterfaceName("nic1")
                                            .withMacAddress("")
                                            .withIpConfigurations(
                                                Arrays
                                                    .asList(
                                                        new NetworkInterfaceIpConfiguration()
                                                            .withIpAllocationMethod(IpAllocationMethod.DYNAMIC)
                                                            .withIpAddress("")
                                                            .withSubnet("")
                                                            .withGateway("")
                                                            .withIpVersion(IpVersion.IPV4)))
                                            .withVmSwitchType(VMSwitchType.MANAGEMENT),
                                        new NetworkInterface()
                                            .withNetworkInterfaceName("nic2")
                                            .withMacAddress("DC-97-F8-79-16-7D")
                                            .withIpConfigurations(
                                                Arrays
                                                    .asList(
                                                        new NetworkInterfaceIpConfiguration()
                                                            .withIpAllocationMethod(IpAllocationMethod.DYNAMIC)
                                                            .withIpAddress("")
                                                            .withSubnet("")
                                                            .withGateway("")
                                                            .withIpVersion(IpVersion.IPV4)))
                                            .withVmSwitchType(VMSwitchType.WAN)))))
            .create();
    }
}
