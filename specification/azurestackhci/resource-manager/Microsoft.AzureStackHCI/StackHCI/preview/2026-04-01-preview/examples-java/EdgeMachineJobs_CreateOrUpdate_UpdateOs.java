
import com.azure.resourcemanager.azurestackhci.models.DeploymentMode;
import com.azure.resourcemanager.azurestackhci.models.IpAddressRange;
import com.azure.resourcemanager.azurestackhci.models.IpAssignmentType;
import com.azure.resourcemanager.azurestackhci.models.NetworkAdapter;
import com.azure.resourcemanager.azurestackhci.models.NetworkConfiguration;
import com.azure.resourcemanager.azurestackhci.models.OSOperationType;
import com.azure.resourcemanager.azurestackhci.models.OnboardingConfiguration;
import com.azure.resourcemanager.azurestackhci.models.OnboardingResourceType;
import com.azure.resourcemanager.azurestackhci.models.OsProvisionProfile;
import com.azure.resourcemanager.azurestackhci.models.ProvisionOsJobProperties;
import com.azure.resourcemanager.azurestackhci.models.ProvisioningOsType;
import com.azure.resourcemanager.azurestackhci.models.ProvisioningRequest;
import com.azure.resourcemanager.azurestackhci.models.SecretType;
import com.azure.resourcemanager.azurestackhci.models.TargetDeviceConfiguration;
import com.azure.resourcemanager.azurestackhci.models.TimeConfiguration;
import com.azure.resourcemanager.azurestackhci.models.UserDetails;
import com.azure.resourcemanager.azurestackhci.models.WebProxyConfiguration;
import java.util.Arrays;

/**
 * Samples for EdgeMachineJobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeMachineJobs_CreateOrUpdate_UpdateOs.json
     */
    /**
     * Sample code: EdgeMachineJobs_CreateOrUpdate_UpdateOs.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineJobsCreateOrUpdateUpdateOs(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineJobs().define("UpdateOs").withExistingEdgeMachine("ArcInstance-rg", "machine1")
            .withProperties(
                new ProvisionOsJobProperties().withDeploymentMode(DeploymentMode.DEPLOY)
                    .withProvisioningRequest(new ProvisioningRequest().withTarget(ProvisioningOsType.AZURE_LINUX)
                        .withOsProfile(new OsProvisionProfile().withOsName("AzureLinux").withOsType("AzureLinux")
                            .withOsVersion("3.1").withOsImageLocation("https://aka.ms/aep/azlinux3.1")
                            .withVsrVersion("1.1.0")
                            .withImageHash("sha256:b2c3d4e5f6789012345678901234567890abcdef1234567890abcdef12345678")
                            .withGpgPubKey("fakeTokenPlaceholder").withOperationType(OSOperationType.UPDATE))
                        .withUserDetails(Arrays
                            .asList(new UserDetails().withUserName("edgeuser").withSecretType(SecretType.KEY_VAULT)
                                .withSecretLocation("fakeTokenPlaceholder")
                                .withSshPubKey(
                                    Arrays.asList("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC7... edgeuser@example.com"))))
                        .withOnboardingConfiguration(new OnboardingConfiguration()
                            .withType(OnboardingResourceType.HYBRID_COMPUTE_MACHINE)
                            .withResourceId(
                                "/subscriptions/ff0aa6da-20f8-44fe-9aee-381c8e8a4aeb/resourceGroups/bhukumar-test-rg/providers/Microsoft.HybridCompute/machines/bkumar-t1")
                            .withLocation("eastus").withTenantId(
                                "72f988bf-86f1-41af-91ab-2d7cd011db47")
                            .withArcVirtualMachineId("634b9db8-83e1-46ed-b391-c1614e2d0097"))
                        .withDeviceConfiguration(new TargetDeviceConfiguration()
                            .withNetwork(new NetworkConfiguration().withNetworkAdapters(Arrays.asList(
                                new NetworkAdapter().withIpAssignmentType(IpAssignmentType.AUTOMATIC).withIpAddress("")
                                    .withIpAddressRange(new IpAddressRange().withStartIp("").withEndIp(""))
                                    .withGateway("").withSubnetMask("").withDnsAddressArray(Arrays.asList("8.8.8.8"))
                                    .withVlanId("0"))))
                            .withHostName("634b9db8-83e1-46ed-b391-c1614e2d0097")
                            .withWebProxy(new WebProxyConfiguration().withConnectionUri("https://microsoft.com/a")
                                .withPort("").withBypassList(Arrays.asList()))
                            .withTime(new TimeConfiguration().withPrimaryTimeServer("").withSecondaryTimeServer("")
                                .withTimeZone("UTC")))
                        .withCustomConfiguration(
                            "eyJjdXN0b21Db25maWciOiAiZXhhbXBsZSBiYXNlNjQgZW5jb2RlZCBjb25maWcifQ==")))
            .create();
    }
}
